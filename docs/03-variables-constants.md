# Chapter 03 — Variables & Constants

## Declaring variables

Go provides two main ways to declare a variable:

```go
// 1. Explicit declaration with a type
var age int = 30

// 2. Short declaration — type is inferred from the value
name := "Alice"   // name is a string
```

The short declaration `:=` is the most common form, but it **only works inside functions**. At package level you must use `var` (or `const`).

You can declare several at once:

```go
var a, b int = 1, 2
var x, y = 1, "two"       // types inferred
var c, d, e int           // all default to 0
```

Or group them:

```go
var (
    name string = "Gopher"
    age  int    = 7
    active bool = true
)
```

## Zero values

Every variable has a well-defined default if you do not initialize it:

| Type | Zero value |
|------|-----------|
| `int`, `int8`…`uint64`, `uintptr` | `0` |
| `float32`, `float64` | `0` |
| `complex64`, `complex128` | `0 + 0i` |
| `bool` | `false` |
| `string` | `""` (empty string) |
| pointer, slice, map, chan, func, interface | `nil` |
| struct | all fields zeroed |

```go
var count int        // 0
var text string      // ""
var flag bool        // false
var nums []int       // nil slice
```

There are **no uninitialized variables** in Go — this eliminates a whole class of bugs.

## Reassignment and redeclaration

```go
x := 10
x = 20          // assign a new value

x := 30         // ERROR: no new variables on left side of :=
```

With multiple variables on the left, `:=` can **redeclare** existing variables as long as **at least one variable on the left is new**:

```go
x := 10
y := 20
x, y := 30, 40   // ERROR: both already exist in this scope — no new variable
x, z := 30, 40   // OK: z is new, x is reassigned to 30
fmt.Println(x, z) // 30 40
```

This rule is what makes the `comma ok` idiom (Chapter 09) and error handling (Chapter 13) convenient:

```go
value, err := doSomething()   // value is new, err may be new or existing
```

## Constants

Constants are values that cannot change after declaration, and they must be computable at **compile time**:

```go
const Pi = 3.14159
const Greeting = "Hello"

const (
    A = 1
    B = 2
)
```

Constants can be untyped, which makes them flexible:

```go
const n = 10
var a int = n       // n adapts to int
var b float64 = n   // ...and to float64
```

Typed constants:

```go
const n int64 = 10
```

## `iota` — enumeration made easy

`iota` is a special constant that starts at `0` and increments by 1 for each line in a `const` block:

```go
const (
    Sunday = iota   // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)
```

Skipping values and expressions:

```go
const (
    _  = iota             // 0 — discard
    KB = 1 << (10 * iota) // 1 << 10  = 1024
    MB                    // 1 << 20
    GB                    // 1 << 30
    TB                    // 1 << 40
)
```

> Repeating the previous expression (like `MB` above) is idiomatic Go — it copies `1 << (10 * iota)` with the new `iota`.

## Scope

Go uses lexical (block) scope:

- **Package scope** — declared outside any function; visible throughout the package.
- **Local/function scope** — declared inside `{}`; visible only there.
- **Shadowing** — an inner declaration hides an outer one:

```go
x := 1
{
    x := 2       // shadows outer x
    fmt.Println(x) // 2
}
fmt.Println(x)     // 1
```

Shadowing is legal but should be used sparingly; it can make code harder to read.

## Naming conventions

- Use `camelCase` for variables and functions: ` userName ` → `userName`.
- Use `UPPER_SNAKE_CASE` only for constants you want exported? No — Go uses `CamelCase`, with the first letter capitalized to **export** from a package:
  ```go
  const MaxRetries = 3   // exported
  const maxRetries = 3   // package-private
  ```
- Keep names short in small scopes (`i`, `err`, `ctx`), descriptive in large scopes.
- Avoid underscores in Go names (`user_name` is not idiomatic; use `userName`).

## Runnable example

See `examples/02-variables/main.go`.

## Chapter summary

- `var name Type = value` for explicit declaration; `name := value` for inferred, local declarations.
- Every type has a zero value; no variable is ever uninitialized.
- `:=` requires at least one new variable on the left side when redeclaring.
- `const` holds compile-time values; `iota` generates enumerated constants.
- Capitalize the first letter to export a name from a package.

## Exercises

1. Declare three variables: an `int`, a `string`, and a `bool`, using both `var` and `:=`. Print their zero values.
2. Use `iota` to create constants for the four seasons starting at 1.
3. Show that a constant can be assigned to both an `int` and a `float64`.
4. Demonstrate shadowing in a nested block.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 03)*
