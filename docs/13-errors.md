# Chapter 13 — Error Handling

Go has **no exceptions**. Errors are ordinary values returned from functions, normally as the last return value.

## The error interface

```go
type error interface {
    Error() string
}
```

## Returning and checking errors

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println(result)
```

The pattern `if err != nil { return err }` is ubiquitous. You will write it hundreds of times.

### Creating errors

```go
import "errors"

err := errors.New("something went wrong")

err := fmt.Errorf("failed to open %s: %w", path, err)  // with wrapping
```

Use lowercase error messages without trailing punctuation — they are often concatenated with context.

## Error wrapping (`%w`)

Go 1.13 introduced error wrapping. Use `%w` in `fmt.Errorf` to keep the original error accessible:

```go
if err := readConfig(path); err != nil {
    return fmt.Errorf("read config %q: %w", path, err)
}
```

`%v` formats but does **not** wrap — use `%w` when you want callers to inspect the cause.

### Inspecting errors

```go
if errors.Is(err, os.ErrNotExist) {
    // handle missing file
}

var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("path:", pathErr.Path)
}
```

- `errors.Is(err, target)` — checks the chain for a specific error value.
- `errors.As(err, &target)` — finds the first error in the chain matching a type.

`errors.Join(errs...)` (Go 1.20+) combines multiple errors into one while keeping `Is`/`As` working:

```go
return errors.Join(err1, err2)
```

## Custom error types

Define a type with an `Error() string` method:

```go
type ParseError struct {
    Line int
    Msg  string
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("parse error at line %d: %s", e.Line, e.Msg)
}

// Optional: implement Unwrap for error chains
func (e *ParseError) Unwrap() error { return e.Err }
```

Using a struct gives callers structured information instead of parsing strings.

### Sentinel errors

Package-level error values callers can compare:

```go
var ErrNotFound = errors.New("not found")

func Lookup(k string) (Value, error) {
    // ...
    return Value{}, ErrNotFound
}

if errors.Is(err, ErrNotFound) { ... }
```

Export sentinels with an `Err` prefix.

## Error handling strategies

1. **Handle and recover** at the level that knows what to do:
   ```go
   if err != nil { fallback() }
   ```
2. **Propagate** with added context:
   ```go
   return fmt.Errorf("fetch user %d: %w", id, err)
   ```
3. **Retry** transient failures (with backoff).
4. **Log and continue** for non-fatal errors.
5. **Fail fast** at startup — `log.Fatal` in `main` is acceptable.

> **Important:** do not both log and return the same error. Pick one. Double logging creates confusing duplicate messages. A common rule: "log at the top, wrap as it bubbles."

## `panic` and `recover`

`panic` unwinds the stack, running deferred functions, then terminates the program:

```go
panic("unreachable state")
```

`recover` can stop a panic, but only when called inside a `defer`:

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    panic("boom")
}
```

**Guidelines:**
- Use `panic` only for truly unrecoverable programming bugs (e.g., impossible invariants) or during initialization.
- Library code should return errors, not panic.
- Use `recover` in servers so a single request cannot crash the process (e.g., `net/http` already recovers per handler).

Do **not** use `panic` for normal error handling.

## Deferred cleanup + errors

Remembering to close files is critical; combine `defer` with error handling:

```go
func copyFile(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil { return err }
    defer func() {
        if cerr := out.Close(); err == nil {
            err = cerr
        }
    }()

    _, err = io.Copy(out, in)
    return err
}
```

## Errors in goroutines

Errors cannot be returned from a goroutine's `Goexit`. Send them over a channel or use `errgroup`:

```go
errCh := make(chan error, 1)
go func() {
    errCh <- doWork()
}()
if err := <-errCh; err != nil { ... }
```

## Runnable example

See `examples/12-errors/main.go`.

## Chapter summary

- Errors are values, returned as the last result; check `if err != nil`.
- Wrap with `%w` and inspect with `errors.Is` / `errors.As`.
- `errors.Join` combines multiple errors.
- Custom error types carry structured context.
- Reserve `panic`/`recover` for exceptional cases, not control flow.
- Log once: wrap on the way up, log at the top.

## Exercises

1. Write `ReadFile` that wraps the OS error with the path using `%w`.
2. Define a `ValidationError` type with a field name and message.
3. Use `errors.Is` to check for `os.ErrNotExist` after wrapping.
4. Write a function with `defer`/`recover` that converts a panic into an error.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 13)*
