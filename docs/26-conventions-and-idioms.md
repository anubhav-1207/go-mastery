# Chapter 26 — Conventions & Idioms

## Formatting

- `gofmt` is mandatory. Tabs for indentation, spaces for alignment.
- Lines should generally stay under ~100–120 characters (not enforced).
- Imports in three groups: standard library, third-party, local — separated by blank lines (gofmt keeps group order).

## Naming

| Item | Convention | Example |
|------|-----------|---------|
| Package | lowercase, short, no underscores | `http`, `json`, `httputil` |
| File | lowercase, underscores | `user_service.go` |
| Variable / function | camelCase | `userName`, `getUsers` |
| Exported | CamelCase (capital first) | `GetUsers`, `MaxSize` |
| Constant | CamelCase | `MaxRetries` |
| Interface | `-er` suffix when natural | `Reader`, `Closer`, `Notifier` |
| Acronyms | all caps | `URL`, `ID`, `HTTPServer` |
| Test file | `_test.go` suffix | `user_test.go` |
| Receiver | 1–2 letters, consistent | `func (u *User)` |

**Avoid stutter:** package `http`'s type is `Server`, not `HTTPServer`. Name things relative to their package.

**Avoid underscores in names** except in generated code and test function names (`Test_parseUser`).

## Comments

- Comments on exported identifiers start with the identifier name:
  ```go
  // Parse parses the input and returns a value.
  func Parse(s string) (Value, error) { ... }
  ```
- Use full sentences with punctuation.
- Explain *why*, not *what*, when the code is already clear.
- Package comment above `package`:
  ```go
  // Package mathutil provides simple arithmetic helpers.
  package mathutil
  ```

## Common idioms

### Comma-ok

```go
v, ok := m[key]
v, ok := x.(Type)
v, ok := <-ch
```

### The zero value is meaningful

```go
var buf bytes.Buffer       // ready to use
var wg sync.WaitGroup      // ready to use
var s []int                // nil slice works with append/range
```

### Guard clauses

```go
if err != nil { return err }
if x == nil { return nil }
```

### Blank identifier

```go
for _, v := range slice { ... }        // ignore index
_ = doSomething()                      // explicitly ignore a result
var _ Interface = (*Impl)(nil)         // compile-time interface check
```

The last form asserts at compile time that `*Impl` implements `Interface`.

### `defer` right after checking the error

```go
f, err := os.Open(p)
if err != nil { return err }
defer f.Close()
```

### Naked returns only in short functions

```go
func split(s string) (a, b string) {
    a = s[:1]
    b = s[1:]
    return
}
```

### Constructor naming

`NewX` returns `*X` (or `X`); `New` in a package returns the main type.

```go
func NewClient(opts ...Option) *Client { ... }
```

### Functional options (for configurable constructors)

```go
type Option func(*Client)

func WithTimeout(d time.Duration) Option {
    return func(c *Client) { c.timeout = d }
}

func NewClient(opts ...Option) *Client {
    c := &Client{timeout: 5 * time.Second}
    for _, o := range opts { o(c) }
    return c
}
```

### Error variables

Export sentinel errors as `ErrX`: `var ErrNotFound = errors.New("not found")`.

### Type aliases for clarity

```go
type UserID int64
```

## Package organization idioms

- One responsibility per package.
- Avoid packages named `util`, `common`, `misc` — they become dumping grounds.
- Keep `internal/` for private implementation.
- Expose the smallest API possible.

## API design idioms

- Make the zero value useful.
- Return concrete types; accept interfaces.
- Use options for optional parameters instead of many booleans (`func New(x, bool, bool, bool)` is unreadable).
- Group related functionality into methods.

## Code review signals

| Smell | Fix |
|-------|-----|
| `if err != nil { panic(err) }` in library | return the error |
| Ignored error `_ = f()` | handle it or comment why |
| Huge `switch` on a type field | use interfaces/polymorphism |
| `map[string]any` everywhere | define structs |
| Global mutable variable | inject dependency |
| Goroutine with no stop condition | add context |

## Runnable example

See `examples/25-idioms/main.go`.

## Exercises

1. Rewrite `HTTPHttpServer` as an idiomatic name.
2. Convert a constructor with three boolean flags to functional options.
3. Add compile-time interface checks for two implementations.
4. Turn a `switch` on a `type` string field into an interface-based design.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 26)*
