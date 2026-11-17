# Chapter 12 — Interfaces

An **interface** is a set of method signatures. A type satisfies an interface **implicitly** — there is no `implements` keyword. If a type has all the required methods, it satisfies the interface.

## Defining and satisfying

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}
func (Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (Cat) Speak() string { return "Meow!" }

func makeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

makeItSpeak(Dog{})  // Woof!
makeItSpeak(Cat{})  // Meow!
```

`Dog` and `Cat` automatically satisfy `Speaker`. This is **structural typing**.

## The empty interface `any`

```go
var x any = 42
x = "hello"
x = []int{1, 2}
```

`any` (alias for `interface{}`) has zero methods, so **every type** satisfies it. Use it sparingly — it discards type safety. Prefer generics (Chapter 19) when you need type-parameterized code.

## Type assertions

Recover the concrete type from an interface:

```go
var x any = "hello"
s := x.(string)          // panics if x is not a string
s, ok := x.(string)      // safe: ok is false on mismatch
```

For interfaces with methods:

```go
var r io.Reader = os.Stdin
f, ok := r.(*os.File)    // check if it is a *os.File
```

## Type switches

```go
func describe(v any) {
    switch x := v.(type) {
    case int:
        fmt.Println("int:", x*2)
    case string:
        fmt.Println("string:", strings.ToUpper(x))
    case nil:
        fmt.Println("nil")
    default:
        fmt.Printf("unknown type %T\n", x)
    }
}
```

A type switch is cleaner than a chain of type assertions.

## Interface values are two words

An interface value holds a **dynamic type** and a **dynamic value**:

```go
var s Speaker = Dog{}   // type = Dog, value = Dog{}
```

An interface is `nil` only when **both** parts are nil:

```go
var s Speaker           // nil interface
var p *Dog = nil
s = p                   // interface is NOT nil: type *Dog, value nil
fmt.Println(s == nil)   // false!
s.Speak()               // may panic — *Dog has value nil
```

**Pitfall:** returning a typed nil as an interface. Avoid by returning a literal `nil`:

```go
func newDog() Speaker {
    var d *Dog = nil
    return d             // BAD: non-nil interface
    // return nil         // GOOD
}
```

## The `error` interface

```go
type error interface {
    Error() string
}
```

Any type with an `Error() string` method is an error. This tiny interface is the backbone of Go error handling (Chapter 13).

```go
fmt.Println(err.Error())  // explicit
fmt.Println(err)          // fmt calls Error() automatically
```

## Important standard-library interfaces

| Interface | Methods | Purpose |
|-----------|---------|---------|
| `fmt.Stringer` | `String() string` | custom formatting |
| `error` | `Error() string` | errors |
| `io.Reader` | `Read(p []byte) (n int, err error)` | read bytes |
| `io.Writer` | `Write(p []byte) (n int, err error)` | write bytes |
| `io.Closer` | `Close() error` | release resources |
| `sort.Interface` | `Len`, `Less`, `Swap` | sorting (use `slices.Sort` nowadays) |
| `encoding/json.Marshaler` | `MarshalJSON() ([]byte, error)` | custom JSON |

Compose small interfaces with embedding:

```go
type ReadWriteCloser interface {
    io.Reader
    io.Writer
    io.Closer
}
```

## Accept interfaces, return structs

Idiomatic API design:

```go
// GOOD: accept the narrowest interface you actually need
func Process(r io.Reader) error { ... }

// GOOD: return concrete types so callers get all methods
func NewClient() *Client { ... }
```

Why? Callers can pass anything satisfying the interface, and they receive a concrete value with full capability.

## Interface segregation (small interfaces)

Prefer small, focused interfaces over large ones:

```go
// GOOD
type Reader interface { Read([]byte) (int, error) }

// BAD — forces implementers to write unused methods
type FileLike interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
    Close() error
    Seek(int64, int) (int64, error)
    // ...
}
```

The standard library lives by this: `io.Reader` and `io.Writer` are each one method and compose endlessly.

## Implementations across packages

One motivating use case: consumers define the interface, not implementers.

```go
// In package myapp — define what you need
type Notifier interface { Notify(string) error }

func Alert(n Notifier, msg string) { n.Notify(msg) }

// Any external type with Notify() works, without importing myapp.
```

## Runnable example

See `examples/11-interfaces/main.go`.

## Chapter summary

- Interfaces are satisfied implicitly by having the required methods.
- `any` accepts everything but loses type safety.
- Use `value, ok := x.(T)` and type switches for safe inspection.
- An interface holding a typed nil is **not** nil — beware.
- `error`, `io.Reader`, `io.Writer` are central interfaces to know.
- Prefer small interfaces; accept interfaces, return concrete types.

## Exercises

1. Define a `Shape` interface with `Area() float64` and implement it for `Circle` and `Square`.
2. Write a type switch that handles `int`, `float64`, and `string`.
3. Demonstrate the typed-nil-in-interface pitfall.
4. Define an interface with embedded `io.Reader` and `io.Writer`.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 12)*
