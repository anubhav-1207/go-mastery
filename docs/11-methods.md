# Chapter 11 — Methods

A **method** is a function associated with a specific type. It is declared with a **receiver** — an extra first parameter before the name.

## Declaring methods

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

Call methods with dot syntax:

```go
r := Rectangle{3, 4}
fmt.Println(r.Area())   // 12
r.Scale(2)
fmt.Println(r.Area())   // 48
```

Go automatically takes the address when calling a pointer-receiver method on an addressable value: `r.Scale(2)` is shorthand for `(&r).Scale(2)`.

## Pointer receivers vs value receivers

| Receiver | Copies receiver? | Can mutate? | When to use |
|----------|------------------|-------------|-------------|
| Value `(r T)` | Yes | No | Small, immutable types; read-only methods |
| Pointer `(r *T)` | No | Yes | Need mutation; large structs; must be consistent |

**Key rule:** if **any** method of a type has a pointer receiver, use pointer receivers for **all** its methods. Mixing leads to surprising method sets.

```go
type Counter struct{ n int }

func (c Counter) Value() int { return c.n }       // OK if no pointer methods
func (c *Counter) Inc()      { c.n++ }            // now prefer pointer for Value too
```

Why? A `*Counter` has both methods, but a `Counter` value does **not** have `Inc`. Mixing means some callers cannot call all methods.

**Exceptions:** small immutable types (like `time.Time`) use value receivers throughout. Slices, maps, and channels can be mutated with value receivers because they are descriptors referencing shared data:

```go
type Stack []int
func (s Stack) Push(v int) Stack { return append(s, v) }  // must return because append may reallocate
```

## Method sets

- The **method set of `T`** = methods with receiver `T`.
- The **method set of `*T`** = methods with receiver `T` **plus** methods with receiver `*T`.

This matters for interfaces (Chapter 12): a `*T` satisfies more interfaces than `T`.

## Nil receivers

A pointer receiver can be `nil` — useful for optional behavior:

```go
type Node struct {
    Value int
    Next  *Node
}

func (n *Node) String() string {
    if n == nil {
        return "<nil>"
    }
    return strconv.Itoa(n.Value)
}
```

## Embedding and method promotion

Methods of embedded types are promoted:

```go
type Reader struct{}
func (Reader) Read() string { return "data" }

type File struct {
    Reader            // embedded
    Name string
}

f := File{}
fmt.Println(f.Read())  // promoted method → "data"
```

Embedding a **pointer** promotes methods of both `T` and `*T`:

```go
type File struct {
    *Reader
}
```

## Methods on non-struct types

You can define methods on any **named local type**:

```go
type MyInt int
func (m MyInt) Double() MyInt { return m * 2 }

type Vec []float64
func (v Vec) Sum() float64 {
    total := 0.0
    for _, x := range v {
        total += x
    }
    return total
}
```

You cannot define methods on a type you do not own (e.g., you cannot add a method to `int` or to a type from another package).

## Stringer and formatting

Implementing `String() string` makes your type work with `fmt.Println` and `%v`:

```go
func (p Person) String() string {
    return fmt.Sprintf("Person(%s, %d)", p.Name, p.Age)
}
```

This is the `fmt.Stringer` interface. Similarly, `Error() string` makes a type an `error` (Chapter 13).

## Constructors + methods

```go
func NewRectangle(w, h float64) *Rectangle {
    if w <= 0 || h <= 0 {
        return nil
    }
    return &Rectangle{Width: w, Height: h}
}
```

## Runnable example

See `examples/10-methods/main.go`.

## Chapter summary

- Methods have a receiver: `func (r T) Method()` or `func (r *T) Method()`.
- Use pointer receivers to mutate, avoid large copies, and stay consistent.
- `T` and `*T` have different method sets; `*T` is a superset.
- Methods work on any named local type, not just structs.
- `String() string` integrates with `fmt`; `Error() string` integrates with errors.

## Exercises

1. Add an `Area()` and `Perimeter()` method to `Rectangle`.
2. Write a `BankAccount` with `Deposit` and `Withdraw` methods using a pointer receiver.
3. Demonstrate that a value of type `T` cannot call a pointer-receiver method when it is not addressable.
4. Implement `String()` for a custom type and print it with `%v`.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 11)*
