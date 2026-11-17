# Project 05 — URL Shortener

**Skills:** `net/http`, pattern routing, middleware, JSON APIs, `sync.RWMutex`, `httptest`.

## Run

```bash
go run ./projects/05-url-shortener
curl -X POST localhost:8080/api/shorten -d '{"url":"https://go.dev"}'
# {"code":"a","url":"https://go.dev"}
curl -i localhost:8080/a     # 302 redirect to https://go.dev
curl localhost:8080/api/stats
```

## Test

```bash
go test -race ./projects/05-url-shortener
```

## What to notice

- The store guards shared state with an `RWMutex` — safe under concurrent requests.
- Handlers are installed with Go 1.22 method+path patterns and `r.PathValue`.
- Middleware (`requestID`) wraps the mux without changing handlers.
- Tests use `httptest` — no real server or port needed.
