# Chapter 14 — Pointers

A **pointer** stores the memory address of another value. Go has pointers but no pointer arithmetic (unlike C), which removes a large class of bugs.

## Address and dereference

```go
x := 42
p := &x          // p is *int, "address of x"

fmt.Println(p)   // 0xc000... (an address)
fmt.Println(*p)  // 42 — dereference: "value at p"

*p = 100
fmt.Println(x)   // 100 — modified through the pointer
```

| Syntax | Meaning |
|--------|---------|
| `&x` | address of `x` |
| `*p` | value pointed to by `p` |
| `*int` | type "pointer to int" |

## Zero value

The zero value of any pointer type is `nil`:

```go
var p *int    // nil
fmt.Println(p == nil)  // true
// fmt.Println(*p)     // panic: nil pointer dereference
```

Always check before dereferencing when a pointer may be nil:

```go
if p != nil {
    fmt.Println(*p)
}
```

## Pointers to structs

Very common:

```go
type Person struct{ Name string }

p := &Person{Name: "Alice"}
fmt.Println(p.Name)     // sugar for (*p).Name
p.Name = "Bob"          // sugar for (*p).Name = "Bob"
```

`new(T)` allocates a zeroed value and returns `*T`:

```go
p := new(Person)   // same as &Person{}
```

## Passing by value vs by pointer

```go
func byValue(n int)    { n = 99 }      // caller's n unchanged
func byPointer(n *int) { *n = 99 }     // caller's n changed

x := 1
byValue(x)
fmt.Println(x)  // 1
byPointer(&x)
fmt.Println(x)  // 99
```

Everything in Go is passed **by value** — a pointer is itself passed by value, but it points to shared data.

When to pass pointers:
- You need to modify the argument.
- The value is a large struct (avoid copying).
- The type must not be copied (e.g., `sync.Mutex`, `os.File`, `bytes.Buffer`).
- You need to represent "absent" (nil pointer).

When to pass values:
- Small types (`int`, `bool`, small structs).
- You want immutability / no aliasing surprises.

## Escape analysis

Go decides at compile time whether a value can live on the stack or must move to the heap. If a local variable's address escapes the function (e.g., returned as a pointer), it is heap-allocated and garbage-collected.

```go
func makeCounter() *int {
    n := 0         // escapes to heap
    return &n
}
```

Check the compiler's decision:

```bash
go build -gcflags='-m' .
# ./main.go:3:9: &n escapes to heap
```

**Takeaway:** don't fear pointers or heap allocation prematurely — write clear code and profile (Chapter 24). Escapes are normal and cheap in Go.

## Pointer aliasing

Two pointers to the same value:

```go
a := 10
p, q := &a, &a
*q = 20
fmt.Println(*p)  // 20 — both see the change
```

Aliasing is powerful but can cause subtle bugs when shared data is mutated from multiple places (especially concurrently — synchronize!).

## Slices and maps are already descriptors

A slice or map value contains a pointer internally, so you often do **not** need a pointer to it to modify its elements:

```go
func add(s []int, m map[string]int) {
    s[0] = 99            // affects caller
    m["x"] = 1           // affects caller
}
```

You **do** need a pointer if you want to reassign the descriptor itself:

```go
func grow(s *[]int) {
    *s = append(*s, 1, 2, 3)  // caller's slice header updated
}
```

Alternatively, return the new slice (more idiomatic):

```go
func grow(s []int) []int { return append(s, 1, 2, 3) }
```

## Receivers revisited

Pointer receivers on methods (Chapter 11) are just pointers used for mutation and avoiding copies.

## Unsafe pointers (avoid)

`unsafe.Pointer` and `uintptr` bypass Go's safety guarantees. They are needed only for system-level programming and may break with future Go versions. Treat them as a last resort.

## Runnable example

See `examples/13-pointers/main.go`.

## Chapter summary

- `&x` takes an address; `*p` dereferences.
- Pointer zero value is `nil`; dereferencing nil panics.
- Go passes everything by value; pointers allow sharing/mutation.
- Use pointers for large structs, mutation, or "absent" semantics.
- Escape analysis decides stack vs heap; trust it until profiling says otherwise.
- Slices/maps are descriptors — you rarely need pointers to them unless reassigning.

## Exercises

1. Write `swap(a, b *int)` that exchanges two integers.
2. Demonstrate that a pointer to a struct lets you modify its fields in a function.
3. Use `go build -gcflags='-m'` to find an escaping variable.
4. Write a function that returns a pointer to a local variable and verify it still works.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 14)*
