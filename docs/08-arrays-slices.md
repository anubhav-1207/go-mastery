# Chapter 08 — Arrays & Slices

## Arrays

An array is a **fixed-length** sequence of a single type. Its length is part of its type:

```go
var a [3]int              // [0 0 0]
b := [3]int{1, 2, 3}      // [1 2 3]
c := [...]int{1, 2, 3, 4} // compiler counts 4 → [4]int
d := [5]int{2: 42}        // [0 0 42 0 0]
```

```go
fmt.Println(len(b))        // 3
b[0] = 10                  // indexing is 0-based
```

Arrays are **values**: assigning one copies the whole array.

```go
x := [2]int{1, 2}
y := x       // y is a copy
y[0] = 99
fmt.Println(x) // [1 2] — unchanged
```

Passing an array to a function also copies it (use a pointer or a slice to avoid that):

```go
func modify(a [3]int) { a[0] = 99 }    // no effect on caller
func modifyPtr(a *[3]int) { a[0] = 99 } // affects caller
```

Arrays are relatively rare in application code; slices are the workhorse.

## Slices

A slice is a **descriptor** of a contiguous segment of an array: it holds a pointer, a length, and a capacity.

```go
s := []int{1, 2, 3}       // slice literal (no length)
var n []int               // nil slice: len 0, cap 0
```

Creating a slice from an array:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]    // [2 3 4] — len 3, cap 4 (from index 1 to end)
s = arr[:3]      // [1 2 3]
s = arr[2:]      // [3 4 5]
s = arr[:]       // whole array
```

Full slice expression limits capacity (useful to avoid overwriting hidden data):

```go
s := arr[1:3:3]  // len 2, cap 2
```

### `len` and `cap`

```go
s := make([]int, 3, 5)
fmt.Println(len(s), cap(s))  // 3 5
```

- `len(s)` — number of accessible elements.
- `cap(s)` — number of elements from the start of the underlying array to its end.

### `make`

```go
s := make([]int, 3)      // len 3, cap 3 → [0 0 0]
s := make([]int, 0, 10)  // len 0, cap 10 → preallocate
```

Preallocating capacity when you know (or estimate) the size avoids repeated reallocations:

```go
result := make([]int, 0, len(input))
for _, v := range input {
    result = append(result, transform(v))
}
```

### `append`

```go
s := []int{1, 2}
s = append(s, 3)         // [1 2 3]
s = append(s, 4, 5)      // [1 2 3 4 5]
s = append(s, other...)  // spread
```

`append` may allocate a new underlying array if capacity is insufficient, so **always assign the result**.

Growth strategy: Go roughly doubles capacity (with optimizations for small sizes). You cannot rely on exact growth.

### `copy`

```go
dst := make([]int, len(src))
n := copy(dst, src)   // copies min(len(dst), len(src)) elements
```

`copy` is useful for resizing without aliases:

```go
s = s[:len(s)+1]      // extend if capacity allows
copy(s[i+1:], s[i:])  // shift elements right
s[i] = value          // insert
```

### Slices are references to arrays

Unlike arrays, assigning a slice makes both variables share the same backing array:

```go
a := []int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a) // [99 2 3] — a changed!
```

This is usually what you want, but be aware when modifying slices inside functions.

### `nil` vs empty slice

```go
var a []int        // nil slice
b := []int{}       // empty, non-nil slice
c := make([]int,0) // empty, non-nil slice

fmt.Println(a == nil, b == nil) // true false
```

Both have `len == 0` and behave the same with `range` and `append`. JSON marshals `nil` as `null` but empty as `[]` — a common gotcha.

### Multi-dimensional slices

```go
grid := make([][]int, 3)
for i := range grid {
    grid[i] = make([]int, 4)
}
grid[0][1] = 5
```

### `slices` package (Go 1.21+)

```go
import "slices"

slices.Contains([]int{1,2,3}, 2)         // true
slices.Index([]string{"a","b"}, "b")     // 1
slices.Sort([]int{3,1,2})                // in-place → [1 2 3]
slices.SortFunc(users, func(a,b User) int { return cmp.Compare(a.Age, b.Age) })
slices.Clone(s)                          // copy
slices.Equal(a, b)                       // deep equality
slices.Compact(s)                        // remove adjacent duplicates (sorted input)
```

Helper functions `slices.Map`, `slices.Filter`, `slices.Reduce` exist in newer versions.

## Strings vs slices

Strings are immutable; converting to `[]byte` copies:

```go
b := []byte("hello")   // new byte slice
b[0] = 'H'
s := string(b)         // "Hello"
```

For in-place mutation of bytes, mutate the `[]byte` then convert back.

## Common pitfalls

1. **Holding references to large arrays** — a small slice of a huge array keeps the whole array alive until the slice is gone. `copy` into a fresh slice to release it.
2. **Assuming `append` never reallocates** — always use the returned slice.
3. **Modifying a slice while ranging** — `range` evaluates length once; appends may not be visited.

## Runnable example

See `examples/07-arrays-slices/main.go`.

## Chapter summary

- Arrays are fixed-size values; slices are dynamic-length descriptors over arrays.
- `make([]T, len, cap)` preallocates capacity for performance.
- `append` may reallocate — always assign its result.
- Slices share their backing array; use `copy`/`slices.Clone` to detach.
- `nil` and empty slices differ in JSON (`null` vs `[]`) but are otherwise similar.

## Exercises

1. Write a function that removes an element at index `i` from a slice, preserving order.
2. Demonstrate that two slices sharing a backing array affect each other.
3. Use `slices.SortFunc` to sort structs by a field.
4. Show the difference between an empty and a nil slice when marshaled to JSON.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 08)*
