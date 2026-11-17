# Go Mastery — Solutions

Detailed answers to [QUESTIONS.md](QUESTIONS.md). Work through the problems first; use this to check and deepen understanding.

---

## Chapter 01 — Introduction

**C1.1** Go was created at Google (2007, released 2009) by Griesemer, Pike, and Thompson to address: slow builds of large C++/Java codebases, complex dependency management, and the difficulty of writing correct concurrent programs. Goals: simple language, fast compilation, built-in concurrency, and a strong standard library.

**C1.2** Composition means assembling behavior from small independent pieces. Structs allow embedding (`type Dog struct { Animal }`) so fields/methods are promoted without an inheritance hierarchy. Interfaces are satisfied implicitly, so a type becomes usable wherever its behavior fits — no declaration of intent couples the types. This avoids the fragile base-class problem and deep hierarchies.

**C1.3** Strong: backend APIs/microservices, CLI tools, cloud infrastructure, concurrent data processing. Weak: GUI desktop apps (few mature toolkits), hard real-time systems (GC pauses), GPU-heavy numerical work (limited ecosystems).

**P1.4** `go version` prints e.g. `go version go1.27.4 linux/arm64`. `go env GOROOT` shows the Go installation; `GOPATH` shows the workspace (default `~/go`).

**P1.5** `fmt.Println` has signature `func Println(a ...any) (n int, err error)`. `...any` is a variadic parameter: it accepts zero or more values of any type. Inside the function, `a` is a `[]any`.

---

## Chapter 02 — Hello, World

**C2.1** `package main` marks an executable; the compiler looks for `func main()` in it. A library package with `func main()` simply has an ordinary function named `main` — it is not an entry point and cannot be run directly.

**C2.2** Go's designers favored fast feedback and clean code: an unused import or variable is almost always a mistake, and the compiler can detect it cheaply. Treating them as errors keeps code tidy without a separate linter and avoids ambiguity about intent.

**C2.3** Go inserts a semicolon at line ends when the last token can end a statement. `func main()` can end a statement, so a newline after it would insert `;`, producing a syntax error before `{`. Hence the brace must be on the same line.

**P2.4**
```go
fmt.Printf("Name: %s, Age: %d, Language: %s\n", "Alice", 30, "Go")
```

**P2.5** `import "os"` unused → `error: "os" imported and not used`. Go 1.22+: `./main.go:5:8: "os" imported and not used`.

---

## Chapter 03 — Variables & Constants

**C3.1** `0`, `0.0`, `false`, `""`, `nil`, `nil`, `nil`, zero value for each field.

**C3.2** At least one variable on the left must be new; existing ones are reassigned. `x, y := 1, 2` fails if both exist in the same scope; `x, z := 1, 2` succeeds when `z` is new.

**C3.3** `iota` is an auto-incrementing constant index starting at 0 per `const` block. `_ = iota` consumes 0; on the next line `iota == 1`, so `KB = 1 << 10 = 1024`.

**P3.4**
```go
var a int; var s string; var ok bool
fmt.Println(a, s, ok) // 0 "" false
outer := 1
{ outer := 2; fmt.Println(outer) } // 2
fmt.Println(outer) // 1
```

**P3.5**
```go
const ( Spring = iota + 1; Summer; Autumn; Winter )
const n = 10
var i int = n
var f float64 = n
```

---

## Chapter 04 — Data Types

**C4.1** `int` and `float64` are the defaults with native machine precision. Use `int64` for ranges beyond 32-bit or binary formats; `uint8`/`byte` for byte data and bit masks.

**C4.2** `byte` = `uint8` (a byte of data). `rune` = `int32` (a Unicode code point). `string` is an immutable byte sequence. `"世界"` is two characters but each CJK rune encodes to 3 UTF-8 bytes, so `len` returns 6.

**C4.3** `type MyInt int` defines a new distinct type with its own method set; conversion is required. `type MyInt = int` is an alias — identical type, no conversion needed.

**P4.4**
```go
s := "Go 语言"
fmt.Println(len(s), len([]rune(s))) // 10 5 (G,o,space = 3 bytes; each CJK = 3 bytes)
for i, r := range s { fmt.Printf("%d %c U+%04X\n", i, r, r) }
```

**P4.5**
```go
r := []rune("hello 世界")
for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }
fmt.Println(string(r))
n := 300
fmt.Println(uint8(n)) // 44 — 300 mod 256 = 44 (wrap-around)
```

---

## Chapter 05 — Operators

**C5.1** Integer `/` truncates toward zero. `%` keeps the sign of the dividend: `-7 % 2 == -1` because `-7 = (-3)*2 + (-1)`.

**C5.2** `++`/`--` are statements so expressions cannot have side effects hidden inside them. `j := i++` is illegal. This prevents confusing constructs like `f(i++)`.

**C5.3** `if x != 0 && 10/x > 2` — the second operand is only evaluated if `x != 0`, avoiding division by zero. Similarly `p != nil && p.Field > 0` avoids a nil dereference.

**P5.4**
```go
const ( Read = 1 << iota; Write; Execute )
perm := Read | Write
fmt.Println(perm&Read != 0, perm&Execute != 0) // true false
```

**P5.5**
```go
var sb strings.Builder
for i := 0; i < 1000; i++ { sb.WriteString("x") }
s := sb.String()
```
`+` allocates a new string each iteration (O(n²) copying); `Builder` grows an internal buffer (amortized O(n)).

---

## Chapter 06 — Control Flow

**C6.1** `for i:=0; i<n; i++`, `for condition`, `for {}`, `for k, v := range x`.

**C6.2** Go randomizes map iteration to prevent reliance on unspecified ordering and to avoid hash-table implementation details leaking into programs. Collect keys, sort them (e.g. `slices.Sort`), then iterate the sorted keys and look up values.

**C6.3** (1) Deferred calls run LIFO. (2) Arguments are evaluated immediately at the `defer` statement. (3) They run when the function returns, including during panic unwinding.

**P6.4**
```go
Loop:
for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
        if i*j > 10 { break Loop }
    }
}
```

**P6.5**
```go
switch {
case n < 0: fmt.Println("negative")
case n == 0: fmt.Println("zero")
case n <= 10: fmt.Println("small")
default: fmt.Println("large")
}
```

---

## Chapter 07 — Functions

**C7.1** Multiple returns make failure explicit in the type signature, eliminate hidden control flow, and avoid exception cost/complexity. `(value, error)` forces the caller to acknowledge failure.

**C7.2** A closure is a function value that captures variables from its enclosing scope. It prints `1 2 3`: `n` lives on as long as the returned function; each call increments the same captured `n`.

**C7.3** Naked returns are fine in very short functions where the named results are obvious. Avoid them in long functions because it becomes unclear what is returned and where.

**P7.4**
```go
func average(xs ...float64) float64 {
    if len(xs) == 0 { return 0 }
    var sum float64
    for _, x := range xs { sum += x }
    return sum / float64(len(xs))
}
```

**P7.5**
```go
func minMax(xs []int) (int, int) {
    min, max := xs[0], xs[0]
    for _, x := range xs[1:] {
        if x < min { min = x }
        if x > max { max = x }
    }
    return min, max
}
func apply(fn func(int) int, xs []int) []int {
    out := make([]int, len(xs))
    for i, x := range xs { out[i] = fn(x) }
    return out
}
```

---

## Chapter 08 — Arrays & Slices

**C8.1** An array has fixed length part of its type and is copied on assignment. A slice is a small descriptor (pointer, len, cap) that references an underlying array; assignment copies the descriptor, so both share data.

**C8.2** `len` = accessible elements, `cap` = elements from the start to the end of the backing array. `append` reallocates when `len+1 > cap`. You must assign the result because reallocation produces a new descriptor/array.

**C8.3** `nil` slice: no backing array, `len`/`cap` 0, equal to nil, marshals to `null`. Empty slice: non-nil, `len` 0, marshals to `[]`. Both work with `range` and `append`.

**P8.4**
```go
func removeAt(s []int, i int) []int { return append(s[:i], s[i+1:]...) }
a := []int{1,2,3}; b := a; b[0] = 9; fmt.Println(a) // [9 2 3]
```

**P8.5**
```go
slices.SortFunc(people, func(a, b Person) int { return cmp.Compare(a.Age, b.Age) })
fresh := slices.Clone(original)
```

---

## Chapter 09 — Maps

**C9.1** Keys must be comparable because hashing requires equality. Numbers, strings, bools, pointers, channels, arrays, comparable structs, interfaces holding comparable values. Slices/maps/functions are not comparable.

**C9.2** `m[k]` returns the zero value for a missing key, indistinguishable from a stored zero (e.g. `0`). `v, ok := m[k]` tells you whether the key existed.

**C9.3** To make programs independent of hash implementation and to discourage reliance on order. Collect keys into a slice, sort, then look up values.

**P9.4**
```go
freq := map[string]int{}
for _, w := range words { freq[w]++ }

set := map[string]struct{}{}
set["x"] = struct{}{}
_, ok := set["x"]
delete(set, "x")
```

**P9.5**
```go
groups := map[string][]string{}
for _, w := range words {
    k := string(w[0])
    groups[k] = append(groups[k], w)
}
clear(groups)
```

---

## Chapter 10 — Structs

**C10.1** Tags are metadata strings interpreted by reflection-based libraries. `encoding/json` uses them to choose field names, omit empty values, and skip fields.

**C10.2** Embedding places a type as an anonymous field, promoting its fields and methods. Unlike inheritance, there is no subtype relationship or dynamic dispatch — it is explicit composition accessed via the embedded field.

**C10.3** `==` on structs requires all fields comparable; slices are not comparable. Use `reflect.DeepEqual` or `slices.Equal` for such fields.

**P10.4**
```go
type Rectangle struct{ W, H float64 }
func Area(r Rectangle) float64 { return r.W * r.H }
a := Rectangle{2,3}; b := a; b.W = 9; fmt.Println(a.W, b.W) // 2 9 (copy)
p := &a; p.W = 9; fmt.Println(a.W) // 9 (shared)
```

**P10.5**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Token string `json:"-"`
}
```

---

## Chapter 11 — Methods

**C11.1** Value receivers copy the receiver; they suit small immutable types and give value semantics. Pointer receivers avoid copies, allow mutation, and are required for types that must not be copied (mutexes, files).

**C11.2** Because `T`'s method set excludes pointer-receiver methods, mixing means a `T` value cannot call all methods but a `*T` can. Consistency avoids surprising API restrictions.

**C11.3** Method set of `T` = methods with receiver `T`. Method set of `*T` includes those plus methods with receiver `*T`.

**P11.4**
```go
type BankAccount struct{ balance int }
func (a *BankAccount) Deposit(n int)  { a.balance += n }
func (a *BankAccount) Withdraw(n int) error {
    if n > a.balance { return errors.New("insufficient funds") }
    a.balance -= n; return nil
}
func (a BankAccount) Balance() int { return a.balance }
```

**P11.5**
```go
func (p Person) String() string { return fmt.Sprintf("Person(%s,%d)", p.Name, p.Age) }
fmt.Println(p)                      // uses String()
var n *Node
fmt.Println((*Node)(nil).String())  // nil receiver must be handled inside String
```

---

## Chapter 12 — Interfaces

**C12.1** A type satisfies an interface automatically by having the required methods — no `implements` declaration. This decouples the interface definition from implementations and lets consumers define interfaces locally.

**C12.2**
```go
var p *Dog = nil
var s Speaker = p
fmt.Println(s == nil) // false: type is *Dog, value is nil
```
The interface is non-nil because it holds type information; calling a method may panic. Return literal `nil` instead.

**C12.3** Functions should take the narrowest interface they need (so callers can pass anything), but return concrete types (so callers get full functionality and do not need type assertions).

**P12.4**
```go
type Shape interface{ Area() float64 }
type Circle struct{ R float64 }
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
type Square struct{ S float64 }
func (s Square) Area() float64 { return s.S * s.S }
```

**P12.5**
```go
switch v := x.(type) {
case int: fmt.Println(v*2)
case float64: fmt.Println(v)
case string: fmt.Println(strings.ToUpper(v))
case nil: fmt.Println("nil")
default: fmt.Printf("%T\n", v)
}
var _ Shape = Circle{} // compile-time check
```

---

## Chapter 13 — Errors

**C13.1** Errors as values make failure explicit in signatures, keep control flow local, avoid hidden stack unwinding, and compose naturally with ordinary code (no try/catch syntax needed).

**C13.2** `%w` wraps the underlying error so `errors.Is`/`errors.As` can inspect the chain; `%v` only formats it into text, losing identity. `Is` matches a sentinel value; `As` extracts a typed error.

**C13.3** `panic` is acceptable for unrecoverable programmer bugs or initialization that must succeed. Libraries should return errors so callers decide how to handle failure.

**P13.4**
```go
data, err := os.ReadFile(path)
if err != nil { return fmt.Errorf("read %q: %w", path, err) }
if errors.Is(err, os.ErrNotExist) { /* missing */ }
```

**P13.5**
```go
type ValidateError struct{ Field string; Err error }
func (e *ValidateError) Error() string { return "invalid " + e.Field }
func (e *ValidateError) Unwrap() error { return e.Err }
// recover wrapper
func safe(fn func()) (err error) {
    defer func() { if r := recover(); r != nil { err = fmt.Errorf("panic: %v", r) } }()
    fn(); return nil
}
```

---

## Chapter 14 — Pointers

**C14.1** `&x` yields the address of `x`; `*p` dereferences a pointer to access the value; `*int` is the type "pointer to int". Zero value is `nil`.

**C14.2** Escape analysis is a compile-time check deciding whether a value can live on the stack or must be heap-allocated. If a local's address escapes (returned, stored globally), Go puts it on the heap, and the GC keeps it alive — so returning `&n` is safe.

**C14.3** You need a pointer to a slice only when you reassign its header (e.g. `append` must update the caller's length/cap). Modifying elements works with a value because the slice descriptor points to shared data.
```go
func mutate(s []int) { s[0] = 9 }             // works
func grow(s *[]int) { *s = append(*s, 1) }    // needed to persist length change
```

**P14.1**
```go
func swap(a, b *int) { *a, *b = *b, *a }
x, y := 1, 2; p, q := &x, &x; *q = 5; fmt.Println(*p) // 5
```

**P14.2** `grow(s *[]int)` dereferences before append; returning `append(s, x)` is usually more idiomatic and composable.

---

## Chapter 15 — Packages & Modules

**C15.1** Capitalized first letter = exported; lowercase = package-private. Applies to all top-level identifiers and struct fields.

**C15.2** `go.mod` declares the module path, Go version, and dependencies. `go.sum` holds checksums verifying downloaded modules. `go mod tidy` adds missing requirements, removes unused ones, and updates `go.sum`.

**C15.3** `cmd/` holds executable entry points; `internal/` holds packages importable only within the module tree. An import cycle is when packages import each other (directly or transitively) — a compile error; break it by extracting shared code.

**P15.4** `mathutil.Add` works; `mathutil.add` → `undefined: mathutil.add` (if referenced) or simply not accessible.

**P15.5** `func init() { fmt.Println("init") }` runs before `main`. Blank import `_ "pkg"` triggers the package's `init` functions for side effects (e.g. registering a DB driver).

---

## Chapter 16 — Goroutines

**C16.1** A goroutine is a lightweight, runtime-managed thread with a small dynamically growing stack (~2 KB initially vs ~1 MB for OS threads), so thousands are cheap.

**C16.2** G = goroutine, M = machine (OS thread), P = processor (scheduling context, count = `GOMAXPROCS`). The runtime schedules Gs onto Ms via Ps, parking blocked goroutines so threads stay busy.

**C16.3** (1) main exits early → wait with WaitGroup/channel; (2) loop capture → Go 1.22 or pass as argument; (3) data races → mutex/atomics + `-race`; (4) leaks → cancellation via context/close.

**P16.4**
```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1) // before go
    go func(id int) { defer wg.Done(); fmt.Println(id) }(i)
}
wg.Wait()
```

**P16.5** Unsafe `shared++` from many goroutines races. Fix with `mu.Lock()/Unlock()` around the increment or `atomic.Int64`. Run `go run -race .` to see the report.

---

## Chapter 17 — Channels

**C17.1** Unbuffered channels rendezvous (blocking until both sides ready) — use for synchronization. Buffered channels queue up to capacity — use to smooth bursts and decouple producer/consumer.

**C17.2** The sender closes, never the receiver. Sending on a closed channel panics. Receiving from a nil channel blocks forever (useful to disable a `select` case).

**C17.3** Fan-out = multiple workers reading from one channel. Fan-in = collecting results from many channels into one. Backpressure arises because a slow consumer blocks a fast producer once the buffer is full.

**P17.4**
```go
func fib(out chan<- int) {
    defer close(out)
    a, b := 0, 1
    for { out <- a; a, b = b, a+b }
}
// fan-out: for w := 0; w < 3; w++ { go worker(jobs) }
```

**P17.5** Use a `WaitGroup` to close the merged channel after all inputs are drained:
```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int); var wg sync.WaitGroup
    for _, c := range cs { wg.Add(1); go func(c <-chan int){ defer wg.Done(); for v := range c { out <- v } }(c) }
    go func(){ wg.Wait(); close(out) }()
    return out
}
```

---

## Chapter 18 — Concurrency Patterns

**C18.1** `select` blocks until one of its channel cases can proceed; `default` makes it non-blocking; `select {}` blocks forever (keeps a program alive).

**C18.2** `ctx` first is a Go convention that makes cancellation ubiquitous and composable. `cancel()` releases timers/resources associated with the context — omitting it leaks them.

**C18.3** Channel: passing jobs to workers, streaming pipeline stages, broadcasting signals. Mutex: protecting a shared counter, guarding a config map, enforcing invariants across multiple fields.

**P18.4** See `projects/04-concurrency-pipeline`: WaitGroup closes result channel after workers finish; `generate` closes job channel.

**P18.5** Use `select { case <-ctx.Done(): return; case v := <-in: ... }` in stages; `sync.Once.Do(setup)` for initialization.

---

## Chapter 19 — Generics

**C19.1** Generics remove duplication while keeping type safety. Pre-generics: `MaxInt`, `MaxFloat`. Generic: `func Max[T cmp.Ordered](a, b T) T`.

**C19.2** `~int` means any type whose underlying type is `int`; `|` builds a union of allowed types; `comparable` is the built-in constraint for types usable as map keys (`==` supported).

**C19.3** Avoid when a simple interface or `any` suffices, when it hurts readability, or for hypothetical future flexibility.

**P19.1–P19.5** Reference implementations live in `examples/18-generics/main.go` (Map, Filter, Number, Stack). `Min`:
```go
func Min[T cmp.Ordered](a, b T) T { if a < b { return a }; return b }
```
**P19.4** Go does not allow methods to declare their own type parameters because the method set of a type must be fixed at type declaration; only receiver type parameters are permitted.

---

## Chapter 20 — File I/O

**C20.1** `Reader.Read(p []byte) (n, err)` and `Writer.Write(p []byte) (n, err)` are minimal abstractions implemented by files, sockets, buffers, compressors. Code written against them composes with everything.

**C20.2** A read can return bytes and `io.EOF` simultaneously. Process `buf[:n]` first, then check EOF, or you silently drop the final bytes.

**C20.3** `filepath.Join` uses the OS separator and cleans the path; string concatenation with `/` breaks on Windows and mishandles duplicate separators/`..`.

**P20.4**
```go
scanner := bufio.NewScanner(f); for scanner.Scan() { n++ }
io.Copy(dstFile, srcFile) // streams efficiently
```

**P20.5** `bufio.Writer` batches writes into a large buffer, reducing syscalls from 10,000 to a few. Always `Flush()`. `os.Stat` gives `Size()` and `ModTime()`.

---

## Chapter 21 — JSON

**C21.1** `json:"name"` renames the output key; `omitempty` omits zero values; `json:"-"` excludes the field entirely.

**C21.2** Without a pointer, `Unmarshal` receives a copy and cannot modify the original. JSON numbers decode as `float64` unless `Decoder.UseNumber()` is set.

**C21.3** `Encoder`/`Decoder` stream to/from `io.Writer`/`io.Reader`, support `DisallowUnknownFields`, `SetEscapeHTML(false)`, and avoid loading whole documents into memory.

**P21.1–P21.5** See `examples/20-json/main.go`. Safe alternative to `map[string]any`: define a struct with typed fields. `any` loses compile-time checks and makes numbers `float64` (precision loss for large integers).

---

## Chapter 22 — Testing

**C22.1** Table-driven tests express many cases compactly and consistently; `t.Run` gives each case a name, isolated failures, and selective execution with `-run`.

**C22.2** `t.Helper()` marks a function as a helper so failure lines point to the caller. `t.TempDir()` creates an auto-cleaned temp directory. `t.Cleanup(fn)` registers teardown.

**C22.3** `ns/op` = nanoseconds per operation; `B/op` = bytes allocated; `allocs/op` = allocations. `-benchmem` enables the allocation columns.

**P22.1–P22.5** See `examples/21-testing/`. `go test -race -cover ./...` runs with race detection and reports coverage percentages per package; use `-coverprofile` + `go tool cover -html` for a visual report.

---

## Chapter 23 — Web Server

**C23.1** `http.Handler` requires `ServeHTTP(w, r)`. `http.HandlerFunc` is a function type with that method, letting ordinary functions be handlers.

**C23.2** Once `WriteHeader` or the first `Write` occurs, headers are sent; later `Header().Set` calls are ignored.

**C23.3** Middleware wraps a handler to add cross-cutting behavior: logging, auth, CORS, request IDs, recovery, rate limiting, compression.

**P23.1–P23.5** See `examples/22-web-server/main.go` and `projects/05-url-shortener/`. Graceful shutdown: listen for `os.Interrupt`/`SIGTERM`, then call `srv.Shutdown(ctx)` with a timeout.

---

## Chapter 24 — Performance

**C24.1** Write clear code → benchmark/profile → identify the hotspot → optimize → re-measure. Never optimize on intuition.

**C24.2** Escape analysis reports which variables move to the heap (causing allocations). `pprof` shows CPU/memory hotspots by function and line.

**C24.3** Preallocate slices; use `strings.Builder`; reuse buffers via `sync.Pool`; avoid `any`/reflection in hot loops; batch I/O with `bufio`; use pointer receivers for large structs.

**P24.1** Preallocation avoids repeated growth and copying — fewer `allocs/op` and `B/op`.
**P24.2** `go test -cpuprofile=cpu.out` + `go tool pprof cpu.out` → `top`.
**P24.3** `testing.AllocsPerRun(n, fn)` returns the average allocations per call; assert it stays 0.
**P24.4** Hand-written encoding avoids reflection but is more error-prone; measure the real gain.
**P24.5** Lower allocation rate reduces GC work and pause frequency; tuning `GOGC` trades memory for CPU without fixing the root cause.

---

## Chapter 25 — Best Practices

**C25.1** Wrapping adds context at each layer so the top-level message explains the full failure chain. Logging at multiple layers duplicates messages; convention: wrap as it bubbles, log once at the boundary.

**C25.2**
```go
func NewService(s Store) *Service { ... } // Store is an interface; return concrete *Service
```
Callers can inject any implementation, and receive a concrete type with all methods.

**C25.3** A goroutine that blocks forever with no cancellation leaks memory and resources for the process lifetime. Every goroutine needs a stop signal (context, closed channel, or done condition).

**P25.1–P25.4** See `examples/24-best-practices/main.go`. The checklist: `gofmt`, `go vet`, `go test`, `go test -race`, `go mod tidy`, no silent errors, no leaky goroutines.

---

## Chapter 26 — Conventions & Idioms

**C26.1** Packages lowercase; camelCase locals; CamelCase exported; acronyms uppercase (`URL`, `ID`); interface `-er` suffix; short receiver names; `_test.go` test files.

**C26.2** `var _ Interface = (*Impl)(nil)` compiles only if `*Impl` implements `Interface` — a zero-cost compile-time assertion.

**C26.3** Functional options are `func(*T)` values passed to a constructor, allowing optional configuration without long parameter lists or boolean flags:
```go
srv := NewServer(":8080", WithTimeout(time.Second), WithTLS(true))
```

**P26.1–P26.4** See `examples/25-idioms/main.go`. Replace `switch v.(type)` on a string field with an interface method; each concrete type implements its own behavior.

---

## Chapter 27 — When to Use What

**C27.1** Array: fixed 3D vector. Slice: dynamic list of users. Map: lookup by ID.

**C27.2** Channel: passing jobs to workers; streaming pipeline stages; broadcasting shutdown. Mutex: shared counter; config map with many readers; protecting multi-field invariants.

**C27.3** `any`: `json.Unmarshal` into generic structures. Generics: `Map`/`Filter` over typed slices. Interface: notifier with different implementations.

**P27.1–P27.4** Use the decision tables in `docs/27-when-to-use-what.md`. Converting `any` to generics:
```go
// before
func First(xs []any) any { return xs[0] }
// after
func First[T any](xs []T) T { return xs[0] }
```

---

## Chapter 28 — Next Steps

**C28.1** `database/sql` (SQL access), `log/slog` (structured logging), `flag` (CLI flags), `regexp` (regular expressions), `embed` (bundle assets), `text/template` (templates), `go/ast` (source analysis).

**C28.2** `slog` (Go 1.21+) provides structured, leveled logging with handlers for JSON/text output — machine-readable and filterable, unlike unstructured `fmt.Println`.

**P28.1–P28.3** Open-ended design/reading exercises. A blog API might use `cmd/api`, `internal/handler`, `internal/store` with a `Store` interface, `internal/model` types, and `errgroup` for startup tasks. `io.Copy` uses `WriterTo`/`ReaderFrom` fast paths, a pooled 32 KB buffer, and `ReadFrom` optimizations to minimize copies and allocations.

---

## Bonus challenge hints

- **B1** `container/list` for LRU order + `map[K]*list.Element` for O(1) lookup; lock the whole operation.
- **B2** `type Ring[T any] struct { buf []T; head, size int }`.
- **B3** Goroutines + a `map[string]bool` guarded by a mutex; channel of URLs with depth.
- **B4** Token bucket: mutex + `tokens`, `last`, refill rate; check/reserve atomically.
- **B5** Fan-out files to workers, each producing partial maps, fan-in and merge under a mutex.
- **B6** `func Chain(m ...Middleware) Middleware { return func(next http.Handler) http.Handler { for i := len(m)-1; i >= 0; i-- { next = m[i](next) }; return next } }`.
- **B7** `signal.Notify(ch, os.Interrupt, syscall.SIGTERM); <-ch; srv.Shutdown(ctx)`.
- **B8** `runtime.Callers` + `runtime.CallersFrames` to build a stack string in `Error()`.
- **B9** `flag.NewFlagSet` per subcommand with `fs.Parse(args[1:])`.
- **B10** `//go:embed static/*` + `embed.FS` + `http.FileServer(http.FS(fs))`.
- **B11** Fuzz random strings; check against `strconv.Atoi`; save failing inputs to `testdata/fuzz`.
- **B12** Use `go test -bench` and `-cpuprofile`, then show the improved `ns/op` and `allocs/op`.

---

*End of solutions. Revisit any chapter whose answers surprised you, then rebuild the projects from memory.*
