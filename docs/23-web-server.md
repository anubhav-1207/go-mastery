# Chapter 23 — Building a Web Server

Go's `net/http` package is production-grade and remarkably simple.

## A minimal server

```go
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, web!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}
```

Visit <http://localhost:8080/>.

## Handlers

A handler is anything implementing:

```go
type Handler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

`http.HandlerFunc` adapts a plain function:

```go
type HandlerFunc func(http.ResponseWriter, *http.Request)
```

### The request

```go
r.Method            // "GET", "POST", ...
r.URL.Path          // "/users/42"
r.URL.Query()       // ?a=1 → map
r.Header.Get("...") // header value
r.Body              // io.ReadCloser
r.Context()         // cancellation context
```

Reading JSON:

```go
var in Payload
if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
    http.Error(w, "invalid JSON", http.StatusBadRequest)
    return
}
```

### The response

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(result)
```

**Order matters:** set headers *before* `WriteHeader`/`Write`; once bytes are written, headers are locked.

## Serving static files

```go
http.Handle("/static/", http.StripPrefix("/static/",
    http.FileServer(http.Dir("public"))))
```

## RESTful routing (Go 1.22+ pattern matching)

```go
mux := http.NewServeMux()

mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("PUT /users/{id}", updateUser)
mux.HandleFunc("DELETE /users/{id}", deleteUser)

func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    // ...
}
```

Earlier Go versions need a router like `github.com/gorilla/mux` or `chi`.

## Middleware

Middleware wraps handlers:

```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
    })
}

mux := http.NewServeMux()
handler := logging(mux)
http.ListenAndServe(":8080", handler)
```

Common middleware: logging, authentication, CORS, request ID, recovery, rate limiting.

### Recovery middleware

```go
func recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                http.Error(w, "internal error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

## Graceful shutdown

```go
srv := &http.Server{
    Addr:    ":8080",
    Handler: mux,
}

go func() {
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

// on SIGTERM/SIGINT:
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

Always set `ReadHeaderTimeout` on production servers to mitigate Slowloris attacks:

```go
srv := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
}
```

## JSON API example

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    u.ID = 1
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}
```

## Testing handlers

```go
func TestHello(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    hello(rec, req)

    if rec.Code != http.StatusOK { t.Fatal(...) }
    if rec.Body.String() != "Hello, web!\n" { t.Fatal(...) }
}
```

## Runnable example

See `examples/22-web-server/main.go`.

## Chapter summary

- `http.Handler` is the core interface; `HandleFunc` adapts functions.
- Read request data from `r`, write responses to `w` (set headers first).
- Go 1.22 mux supports method + path patterns with `{id}` and `r.PathValue`.
- Middleware = `func(http.Handler) http.Handler`.
- Use graceful shutdown and timeouts in production.

## Exercises

1. Build a server with `GET /hello?name=...` that greets the name.
2. Add a middleware that adds an `X-Request-Id` header.
3. Implement a JSON API with `POST /echo` that returns the request body.
4. Write an `httptest` test for a handler returning 404.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 23)*
