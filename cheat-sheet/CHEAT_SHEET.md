# Go Cheat Sheet

A compact reference. Keep it open while coding.

---

## Program structure

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello")
    os.Exit(0)
}
```

---

## Declarations

```go
var x int = 5        // declaration
var y = 10           // inferred
z := 15              // short declaration (inside functions only)
const Pi = 3.14
const (
    A = iota         // 0
    B                // 1
    C                // 2
)

var (                // grouped
    a int
    b string
)
```

---

## Basic types

| Category | Types |
|----------|-------|
| Signed int | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned | `uint`, `uint8`/`byte`, `uint16`, `uint32`, `uint64`, `uintptr` |
| Float | `float32`, `float64` |
| Complex | `complex64`, `complex128` |
| Other | `bool`, `string`, `rune` (`int32`) |
| Zero values | `0`, `0.0`, `false`, `""`, `nil` |

---

## Operators

```go
+  -  *  /  %           // arithmetic
&  |  ^  &^  <<  >>     // bitwise
+= -= *= /= %=          // compound assignment
&& || !                 // logical (short-circuit)
== != < <= > >=         // comparison
i++  i--                // statements only
```

---

## Control flow

```go
if x > 0 {
} else if x == 0 {
} else {
}

if err := f(); err != nil { return err }   // init statement

for i := 0; i < 10; i++ {}
for condition {}
for {}
for k, v := range collection {}

switch x {
case 1, 2:
    // no fallthrough by default
default:
}

switch {                 // expressionless
case x < 0:
default:
}

defer cleanup()          // runs at function exit (LIFO)
```

---

## Functions

```go
func add(a, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
    if b == 0 { return 0, errors.New("div by zero") }
    return a / b, nil
}

func sum(xs ...int) (total int) {       // variadic + named return
    for _, x := range xs { total += x }
    return
}

f := func(x int) int { return x * 2 }   // anonymous function

type Op func(int, int) int               // function type
```

---

## Arrays & slices

```go
var a [3]int                  // array: fixed size
b := [3]int{1, 2, 3}
c := [...]int{1, 2, 3}        // compiler counts

s := []int{1, 2, 3}           // slice
s := make([]int, 3)           // len 3, cap 3
s := make([]int, 0, 10)       // len 0, cap 10

s = append(s, 4)              // may reallocate — always assign
s = append(s, other...)
copy(dst, src)
len(s); cap(s)

s := arr[1:4]                 // slice of array [low:high]
s := arr[1:4:4]               // full slice (limits cap)

slices.Sort(s)                // Go 1.21+
slices.Contains(s, 3)
slices.Clone(s)
```

---

## Maps

```go
m := map[string]int{"a": 1, "b": 2}
m := make(map[string]int, 100)

v := m["a"]                   // zero value if missing
v, ok := m["a"]               // comma-ok
m["c"] = 3
delete(m, "c")
clear(m)                      // Go 1.21+

for k, v := range m {}        // random order!

set := map[string]struct{}{}  // set
set["x"] = struct{}{}
```

---

## Structs & methods

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

p := Person{Name: "Alice", Age: 30}
p := &Person{Name: "Alice"}
var zero Person                // {"", 0}

type Embedded struct { Person } // embedding

func (p Person) Greet() string {        // value receiver
    return "Hi " + p.Name
}
func (p *Person) Birthday() { p.Age++ } // pointer receiver
```

---

## Interfaces

```go
type Speaker interface { Speak() string }

// implemented implicitly: just have the method

var x any = 42                 // empty interface
v, ok := x.(int)               // type assertion
switch v := x.(type) {         // type switch
case int:
case string:
default:
}

var _ Speaker = (*Dog)(nil)    // compile-time check
```

---

## Errors

```go
import "errors"

err := errors.New("message")
err := fmt.Errorf("context: %w", err)  // wrap

if err != nil { return err }
if errors.Is(err, os.ErrNotExist) {}
var pe *os.PathError
if errors.As(err, &pe) {}

return errors.Join(e1, e2)

panic("unrecoverable")          // avoid for normal errors
defer func() {
    if r := recover(); r != nil {}
}()
```

---

## Pointers

```go
x := 5
p := &x           // address
v := *p           // dereference
*p = 10           // modify through pointer
var nilPtr *int   // nil
```

---

## Goroutines & channels

```go
go f()                         // start goroutine

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
}()
wg.Wait()

ch := make(chan int)           // unbuffered
ch := make(chan int, 10)       // buffered
ch <- 1                        // send
v := <-ch                      // receive
v, ok := <-ch                  // closed?
close(ch)                      // sender closes
for v := range ch {}           // until closed

func f(out chan<- int) {}      // send-only
func f(in <-chan int) {}       // receive-only

select {
case v := <-ch:
case ch <- 1:
case <-time.After(time.Second):
default:
}
```

---

## Context & sync

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
<-ctx.Done()                   // cancellation signal
ctx.Err()

var mu sync.Mutex
mu.Lock(); defer mu.Unlock()

var mu sync.RWMutex
mu.RLock(); defer mu.RUnlock()

var once sync.Once
once.Do(setup)

var n atomic.Int64
n.Add(1); n.Load()
```

---

## Generics

```go
func Max[T cmp.Ordered](a, b T) T {
    if a > b { return a }
    return b
}

type Number interface{ ~int | ~float64 }
func Sum[T Number](xs []T) (total T) {
    for _, x := range xs { total += x }
    return
}

type Stack[T any] struct{ items []T }
```

---

## Files & JSON

```go
data, err := os.ReadFile("f.txt")
os.WriteFile("f.txt", data, 0644)

f, err := os.Open("f.txt")
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() { line := scanner.Text() }

w := bufio.NewWriter(f)
fmt.Fprintln(w, "line")
w.Flush()

json.Marshal(v)                // value -> []byte
json.MarshalIndent(v, "", "  ")
json.Unmarshal(b, &v)          // []byte -> value (pointer!)
json.NewEncoder(w).Encode(v)
json.NewDecoder(r).Decode(&v)

// struct tags
`json:"field"`
`json:"field,omitempty"`
`json:"-"`
```

---

## Testing

```go
func TestX(t *testing.T) {
    if got := f(); got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
t.Run("name", func(t *testing.T) {})
t.Helper(); t.TempDir(); t.Cleanup(fn)

func BenchmarkX(b *testing.B) {
    for i := 0; i < b.N; i++ { f() }
}
func ExampleX() { fmt.Println(f()); /* Output: ... */ }
func FuzzX(f *testing.F) { f.Add("seed"); f.Fuzz(...) }
```

```bash
go test ./...
go test -v -race -cover ./...
go test -bench=. -benchmem
go test -cpuprofile=cpu.out ./...
```

---

## HTTP server

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /users/{id}", handler)
id := r.PathValue("id")
http.ListenAndServe(":8080", mux)

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(v)

// middleware
func mw(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}
```

---

## Go commands

```bash
go run .                 # compile + run
go build -o app .        # build binary
go install               # install to GOPATH/bin
go test ./...            # tests
go vet ./...             # static checks
gofmt -w .               # format
go mod init <path>       # new module
go mod tidy              # sync deps
go get pkg@v1.2.3        # add dependency
go doc fmt.Println       # docs
go env                   # environment
go list -m all           # dependency list
```

---

## Common gotchas

| Gotcha | Fix |
|--------|-----|
| Unused import/variable | compile error — remove or use it |
| `if 1 {}` | conditions must be bool |
| `x := 1; x := 2` | second needs a new variable on left |
| Map read missing key | returns zero value; use `v, ok` |
| Map write to `nil` | panic — use `make` |
| Range over map | random order — sort keys |
| `append` result ignored | always assign `s = append(...)` |
| Slice alias mutation | use `copy`/`slices.Clone` |
| Interface with typed nil | not nil — return literal `nil` |
| `defer f.Close()` error ignored | handle explicitly if important |
| `fmt.Fprintln` before header set | headers already sent |
| Goroutine sees `i == N` | Go 1.22 fixed; else pass as param |
| Sending on closed channel | panic — sender closes only |

---

*Companion to the [Go Mastery guide book](../README.md).*
