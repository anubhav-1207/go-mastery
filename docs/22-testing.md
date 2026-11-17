# Chapter 22 — Testing

Go has testing built into the toolchain. Test files end in `_test.go` and live beside the code they test.

## Your first test

```go
// math_test.go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

Run it:

```bash
go test ./...
go test -v ./...        # verbose
go test -run TestAdd    # run one test
```

Test files can use `package mathutil` (same package, access internals) or `package mathutil_test` (external, tests public API only). Prefer external for integration-style tests.

## Table-driven tests

The idiomatic Go pattern — one test function, many cases:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

`t.Run` creates subtests; you can run one with `go test -run 'TestAdd/negative'`.

## Test helpers

```go
func assertEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()   // reports the caller's line, not this one
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

Use `t.Fatal` / `t.Fatalf` when continuing makes no sense (it calls `runtime.Goexit()`).

## Setup and teardown

```go
func TestMain(m *testing.M) {
    // global setup
    code := m.Run()
    // global teardown
    os.Exit(code)
}
```

Per-test cleanup:

```go
t.Cleanup(func() { os.Remove(tmpFile) })
```

## Temporary files and directories

```go
dir := t.TempDir()                       // auto-removed
f, _ := os.CreateTemp(t.TempDir(), "*.json")
```

## Mocking with interfaces

Design dependencies as interfaces, then inject fakes in tests:

```go
type Store interface { Get(string) (string, error) }

type fakeStore struct{}
func (fakeStore) Get(k string) (string, error) { return "value", nil }

func TestService(t *testing.T) {
    svc := NewService(fakeStore{})
    // ...
}
```

## Benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

```bash
go test -bench=. -benchmem ./...
```

Look at `ns/op`, `B/op` (bytes allocated), and `allocs/op`. Use `-benchmem` to track allocations (Chapter 24).

Prevent the compiler from optimizing away the work:

```go
var result int
func BenchmarkAdd(b *testing.B) {
    var r int
    for i := 0; i < b.N; i++ {
        r = Add(1, 2)
    }
    result = r
}
```

## Examples as documentation

```go
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}
```

`go test` runs these and verifies the output. They appear in godoc.

## Fuzz testing (Go 1.18+)

```go
func FuzzParse(f *testing.F) {
    f.Add("valid input")
    f.Fuzz(func(t *testing.T, s string) {
        _, err := Parse(s)
        if err != nil {
            t.Skip()
        }
        // properties that must hold...
    })
}
```

```bash
go test -fuzz=FuzzParse -fuzztime=30s
```

The fuzzer mutates inputs to find crashes and regression cases (saved in `testdata/fuzz/`).

## Race detection and coverage

```bash
go test -race ./...            # detect data races
go test -cover ./...           # coverage percentage
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out  # visual report
```

Aim for meaningful coverage of critical paths, not a arbitrary percentage.

## Testing HTTP handlers

```go
req := httptest.NewRequest("GET", "/", nil)
w := httptest.NewRecorder()
handler(w, req)

if w.Code != http.StatusOK { ... }
```

## Runnable example

See `examples/21-testing/` (includes tests, benchmarks, and an example function).

## Chapter summary

- Tests live in `_test.go` files; run with `go test ./...`.
- Table-driven tests with `t.Run` are the Go idiom.
- `t.Helper()`, `t.TempDir()`, and `t.Cleanup` make tests clean.
- Benchmark with `go test -bench=. -benchmem`.
- Fuzz testing finds edge cases automatically.
- Always run `-race` in CI.

## Exercises

1. Write a table-driven test for a `Divide` function including an error case.
2. Add a benchmark for `strings.Builder` vs `+` concatenation.
3. Write an `Example` function for a `Reverse` function.
4. Create a fuzz test for a JSON parser wrapper.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 22)*
