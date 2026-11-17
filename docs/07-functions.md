# Chapter 07 — Functions

## Function declaration

```go
func add(a int, b int) int {
    return a + b
}
```

When consecutive parameters share a type, you can shorten:

```go
func add(a, b int) int {
    return a + b
}
```

If a function has no return value, omit the type:

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

## Multiple return values

Go functions can return more than one value — this is the idiomatic replacement for exceptions and output parameters:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

### Named return values

You can name the results; they start at their zero values and can be assigned with a bare `return`:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return   // naked return — returns x and y
}
```

Naked returns are convenient in short functions but **discouraged in long ones** — they hide what is returned. Prefer explicit returns when the function spans many lines.

## Variadic functions

`...T` as the last parameter accepts any number of arguments of type `T`:

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum()             // 0
sum(1, 2, 3)      // 6
values := []int{1, 2, 3}
sum(values...)    // 6 — spread a slice
```

`fmt.Printf` is variadic: `func Printf(format string, a ...any)`.

## Functions are values

```go
f := add
fmt.Println(f(2, 3))  // 5

type Operation func(int, int) int

func apply(op Operation, a, b int) int {
    return op(a, b)
}
```

Function types are reference-like; their zero value is `nil`. Calling a nil function panics.

## Anonymous functions and closures

```go
f := func(x int) int {
    return x * x
}
fmt.Println(f(5))  // 25
```

A **closure** captures variables from its surrounding scope:

```go
func counter() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}

c := counter()
fmt.Println(c(), c(), c())  // 1 2 3
```

Each call to `counter()` gets its own `n`.

### Pitfall: loop variable capture

```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)   // BUG: may print 3 3 3
    }()
}
```

Since Go 1.22, loop variables are per-iteration by default, so this is now safe. On older Go versions, pass `i` as a parameter: `go func(i int){...}(i)`.

## Recursion

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

Go does **not** guarantee tail-call optimization, so deep recursion can grow the stack (Go stacks grow dynamically, so it is less dangerous than in other languages, but still).

## Higher-order functions

```go
func mapInts(nums []int, fn func(int) int) []int {
    out := make([]int, len(nums))
    for i, n := range nums {
        out[i] = fn(n)
    }
    return out
}

doubled := mapInts([]int{1, 2, 3}, func(n int) int { return n * 2 })
```

The standard library's `slices` package provides `Map`, `Filter`, `Reduce`-style helpers in recent Go versions (Chapter 08).

## `defer` with functions

Remember `defer` takes a function call; to defer with arguments, wrap in a closure if you need evaluation at return time:

```go
defer fmt.Println("done")           // arguments evaluated now
defer func() { fmt.Println(x) }()   // x read when deferred call runs
```

## Methods vs functions

A **method** is a function with a receiver (Chapter 11). A plain function is called directly. Use methods when the behavior belongs to a type; use functions for operations on data that do not need a receiver.

## Runnable example

See `examples/06-functions/main.go`.

## Chapter summary

- Parameters of the same type may be grouped: `func f(a, b int)`.
- Multiple returns, especially `(value, error)`, are idiomatic.
- Variadic `...T` accepts zero or more arguments; spread with `slice...`.
- Functions are first-class values and can form closures.
- Go 1.22 gives each loop iteration its own variables; older versions require passing them explicitly.

## Exercises

1. Write a variadic `average(nums ...float64) float64` that returns 0 for no args.
2. Write a closure that produces an "adder" function: `makeAdder(5)(3)` returns 8.
3. Implement Fibonacci recursively and iteratively; compare clarity.
4. Write a function that returns `(min, max)` of a slice of ints.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 07)*
