# Chapter 25 — Best Practices

This chapter collects the habits that separate idiomatic, maintainable Go from code that merely compiles.

## 1. Format with `gofmt` — always

```bash
gofmt -w .
```

Never argue about formatting; `gofmt` is the standard. Integrate it into your editor and CI.

## 2. Let `go vet` and linters help

```bash
go vet ./...
```

Consider `golangci-lint` for a comprehensive set of linters (`errcheck`, `staticcheck`, `ineffassign`, `gosimple`).

## 3. Handle every error

```go
// Bad
f, _ := os.Open(path)   // silently ignoring

// Good
f, err := os.Open(path)
if err != nil {
    return fmt.Errorf("open %s: %w", path, err)
}
```

Ignoring an error must be a deliberate, documented choice (`_ = w.Close()` in cleanup with a comment).

## 4. Wrap errors with context

```go
return fmt.Errorf("fetch user %d: %w", id, err)
```

Add what you know at each layer. Use `errors.Is`/`As` at the boundaries that react.

## 5. Return early

```go
// Good: happy path stays unindented
if err != nil {
    return err
}
// ... rest of function
```

Avoid deep nesting; guard clauses are idiomatic.

## 6. Keep functions small and focused

A function should do one thing and have a clear name. If it needs a paragraph to describe, split it.

## 7. Accept interfaces, return structs

```go
func NewService(store Store) *Service { ... }   // Store is an interface
```

Callers depend on narrow behavior, not concrete implementations.

## 8. Define interfaces where they are used

The consumer, not the producer, should define the interface. This keeps interfaces small and relevant.

## 9. Avoid global state

Globals make testing and concurrency harder. Pass dependencies explicitly. Exceptions: immutable constants, registered drivers, `sync.Once` singletons when truly needed.

## 10. Use `defer` for cleanup

```go
f, err := os.Open(p)
if err != nil { return err }
defer f.Close()
```

## 11. Don't use `panic` for normal errors

Return errors. Reserve `panic` for unrecoverable bugs during startup.

## 12. Use `context` for cancellation

First parameter, always call `cancel()`, check `ctx.Done()` in loops and I/O.

## 13. Avoid data races

Synchronize shared mutable state with mutexes, atomics, or channels. Run `-race` in CI.

## 14. Be careful with goroutine leaks

Every goroutine must have a termination path. Use `context`, `close`, or `select` with `Done()`.

## 15. Prefer composition to embedding for behavior

Embedding is fine for simple promotion, but explicit fields are clearer when you need control.

## 16. Keep the zero value useful

Design types so that `var x T` is usable without explicit initialization where possible:

```go
type Counter struct { n int }        // usable as zero value
type Buf struct { items []int }      // append works on nil slice
```

## 17. Name things clearly

- Short names for small scopes (`i`, `err`, `ctx`, `w`, `r`).
- Descriptive names for larger scopes.
- No stuttering: `http.HTTPServer` → `http.Server`.
- Initialisms in caps: `URL`, `ID`, `HTTP`, `JSON`.

## 18. Write documentation comments for exported identifiers

```go
// Serve starts the HTTP server on the given address.
func Serve(addr string) error { ... }
```

## 19. Test behavior, not implementation

Write table-driven tests, use fakes via interfaces, and avoid testing private details.

## 20. Keep dependencies minimal

The standard library is huge and stable. Add a third-party module only when it clearly earns its place.

## 21. Use `go mod tidy` regularly

Keep `go.mod` and `go.sum` accurate and commit both.

## 22. Be explicit in public APIs

Avoid `any` in public signatures when a concrete type or generic works. Explicit APIs are self-documenting and type-safe.

## 23. Document concurrency guarantees

State whether a type is safe for concurrent use. If not, say so in its doc comment.

## 24. Measure before optimizing

Benchmark and profile. Clear code first; fast code second.

## Quick checklist before committing

- [ ] `gofmt -w .`
- [ ] `go vet ./...`
- [ ] `go test ./...`
- [ ] `go test -race ./...` (for concurrent code)
- [ ] `go mod tidy`
- [ ] No ignored errors without reason
- [ ] No goroutine without cancellation

## Runnable example

See `examples/24-best-practices/main.go` for a small program demonstrating several of these patterns together.

## Exercises

1. Refactor a function with deep nesting into guard clauses.
2. Take code that ignores errors and wrap them with context.
3. Design a service with an interface dependency and a fake for testing.
4. Run the checklist on one of the projects and fix what it flags.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 25)*
