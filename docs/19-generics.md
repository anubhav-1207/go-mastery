# Chapter 19 — Generics

Generics were added in Go 1.18. They let you write functions and types parameterized over other types while retaining compile-time type safety.

## Why generics?

Before generics, you either duplicated code per type or used `any` with runtime type assertions:

```go
// Pre-generics: duplicated
func MaxInt(a, b int) int { if a > b { return a }; return b }
func MaxFloat(a, b float64) float64 { ... }
```

With generics:

```go
func Max[T cmp.Ordered](a, b T) T {
    if a > b { return a }
    return b
}
```

## Type parameters

```go
func Map[T any, U any](in []T, fn func(T) U) []U {
    out := make([]U, len(in))
    for i, v := range in {
        out[i] = fn(v)
    }
    return out
}

nums := Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n) })
```

- `[T any]` declares a type parameter named `T` constrained by `any`.
- Type arguments are usually **inferred** from the call: `Max(3, 5)`.
- Explicit instantiation: `Max[int](3, 5)`.

## Constraints

A **constraint** is an interface that specifies what types are allowed.

```go
type Number interface {
    ~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
    var total T
    for _, x := range xs {
        total += x
    }
    return total
}
```

- `|` is a **union** of types.
- `~int` means "any type whose underlying type is `int`" (includes named types like `type Age int`).
- `comparable` is a built-in constraint for types usable as map keys.

```go
func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target { return true }
    }
    return false
}
```

The standard library provides `cmp.Ordered` (all types supporting `<`, `>`, `==`).

## Generic types

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T) {
    s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    n := len(s.items)
    v := s.items[n-1]
    s.items = s.items[:n-1]
    return v, true
}

s := &Stack[string]{}
s.Push("hello")
```

Methods on generic types repeat the type parameter list on the receiver.

## Generic methods? No.

Go does **not** allow methods to have their own type parameters (only the receiver's). If you need that, use a generic function instead.

## Type inference

The compiler infers type arguments from function arguments:

```go
Max(3, 7)           // T = int
Map([]int{}, fn)    // T inferred from slice, U from fn
```

Inference cannot always infer from constraints alone — sometimes you must specify explicitly.

## Standard library generics

- `slices` (Go 1.21+): `Sort`, `SortFunc`, `Contains`, `Index`, `Clone`, `Equal`, `Compact`, `BinarySearch`.
- `maps` (Go 1.21+): `Keys`, `Values`, `Equal`, `Copy`.
- `cmp`: `Compare`, `Less`, `Or`.
- `sync/atomic`: `atomic.Pointer[T]`, `atomic.Value`-like generic types.

## When to use generics

**Good uses:**
- Generic data structures (`Stack[T]`, `Set[T]`, `Cache[K,V]`).
- Algorithms over collections (`Map`, `Filter`, `Reduce`).
- Removing near-identical code for multiple types.

**Avoid:**
- When a simple `any` + interface suffices (e.g., `func Log(v any)`).
- When it makes code harder to read than a concrete version.
- Over-engineering for hypothetical future types.

> "Generics are a tool, not a goal." — write the simple version first.

## Generics vs interfaces

| Need | Use |
|------|-----|
| Behavior shared across types | Interface |
| Algorithm shared across types with type safety | Generics |
| Heterogeneous values at runtime | `any` |
| Different types, same operations, performance-sensitive | Generics |

## Runnable example

See `examples/18-generics/main.go`.

## Chapter summary

- Generics add type parameters: `func F[T Constraint](x T) T`.
- Constraints are interfaces with type sets (`~int | ~float64`, `comparable`).
- Generic types use `[T any]` in declaration and `*Stack[T]` receivers.
- `slices`, `maps`, `cmp` are generic standard packages.
- Use generics for real abstraction; prefer simple code otherwise.

## Exercises

1. Write a generic `Filter[T any](s []T, pred func(T) bool) []T`.
2. Implement a `Set[T comparable]` with Add/Contains/Remove.
3. Use `cmp.Ordered` to write `Min[T](a, b T) T`.
4. Explain why `func (s *Stack[T]) PopGeneric[U any]()` is illegal.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 19)*
