# Chapter 09 — Maps

## What is a map?

A map is Go's built-in **hash table**: an unordered collection of key/value pairs.

```go
m := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
```

Declaration and creation:

```go
var m1 map[string]int          // nil map — cannot write to it
m2 := make(map[string]int)     // empty, usable
m3 := map[string]int{}         // same as m2
m4 := make(map[string]int, 100) // preallocate for ~100 entries
```

> A `nil` map can be **read** (returns the zero value) but writing to it panics. Use `make` before inserting.

## Reading and writing

```go
m["Charlie"] = 35        // insert or update
age := m["Alice"]        // 30
age = m["Unknown"]       // 0 — the zero value; no error!
```

To distinguish "missing" from a stored zero value, use the **comma-ok idiom**:

```go
age, ok := m["Unknown"]
if ok {
    fmt.Println("age:", age)
} else {
    fmt.Println("not found")
}
```

To test presence without the value:

```go
if _, exists := m[key]; exists {
    // key exists
}
```

## Deleting

```go
delete(m, "Bob")   // no-op if key is absent; returns nothing
```

To clear all entries:

```go
clear(m)           // Go 1.21+ — removes all keys in place
```

Or simply reassign `m = make(map[string]int)` (old map becomes garbage).

## Keys and values

- **Keys** must be **comparable** (usable with `==`): numbers, strings, booleans, pointers, channels, interfaces holding comparable values, arrays, and structs whose fields are all comparable.
- **Slices, maps, and functions cannot be keys.**
- **Values** can be any type, including slices, maps, structs, and functions.

```go
type Point struct{ X, Y int }
m := map[Point]string{{1, 2}: "origin-ish"}  // OK — struct is comparable

bad := map[[]int]string{}  // ERROR: unhashable type []int
```

## Iteration order is unspecified

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

Order is **randomized on purpose** — do not rely on it. For deterministic output, collect keys and sort:

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
slices.Sort(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

The `maps` package (Go 1.21+) provides helpers:

```go
maps.Keys(m)       // iterator (Go 1.23+; returns []K in older versions)
maps.Values(m)
maps.Equal(m1, m2) // deep equality
maps.Copy(dst, src)
```

## Maps are reference types

Assigning a map copies the descriptor, not the data:

```go
a := map[string]int{"x": 1}
b := a
b["x"] = 99
fmt.Println(a["x"]) // 99 — same map
```

Passing a map to a function allows modifications to be seen by the caller (unless you reassign the parameter, which only changes the local descriptor).

## Map of slices and slice of maps

Very common patterns:

```go
// One key → many values
groups := map[string][]string{
    "fruits": {"apple", "banana"},
}
groups["fruits"] = append(groups["fruits"], "cherry")

// A list of records
users := []map[string]any{
    {"name": "Alice", "age": 30},
    {"name": "Bob", "age": 25},
}
```

## `map[string]any` — flexible but risky

`any` (alias for `interface{}`) accepts any value, but you lose type safety and must type-assert:

```go
m := map[string]any{"age": 30, "name": "Alice"}
age := m["age"].(int)     // panics if not an int
age, ok := m["age"].(int) // safe form
```

Prefer typed structs (Chapter 10) for structured data.

## Set implementation

Go has no built-in set; use a map with `bool` (or `struct{}`) values:

```go
set := map[string]bool{}
set["a"] = true
if set["a"] { /* present */ }

// Memory-efficient variant: empty struct uses 0 bytes
set2 := map[string]struct{}{}
set2["a"] = struct{}{}
_, ok := set2["a"]
delete(set2, "a")
```

## Performance notes

- Average `O(1)` lookup/insert; worst case `O(n)` on hash collisions (rare).
- Preallocate with `make(map[K]V, expectedSize)` if you know the size.
- Keys are hashed; large structs as keys cost more to hash.
- Maps grow dynamically and rehash — writing while iterating is allowed but newly inserted keys may or may not be visited.

## Concurrency

Maps are **not safe for concurrent writes**. Reading concurrently is fine if there are no writes. For concurrent access use `sync.RWMutex` or `sync.Map` (Chapter 18).

## Runnable example

See `examples/08-maps/main.go`.

## Chapter summary

- Create maps with `make` or a literal; a `nil` map panics on write.
- Use `value, ok := m[key]` to detect missing keys.
- Iteration order is unspecified — sort keys when order matters.
- Keys must be comparable; values can be anything.
- Use `map[T]struct{}` for sets, `clear` to empty a map.
- Guard concurrent writes with locks.

## Exercises

1. Count word frequencies in a `[]string` using a map.
2. Implement a set with `map[string]struct{}` supporting add, contains, and remove.
3. Group users by their first letter using `map[string][]string`.
4. Show why ranging over a map gives different orders across runs.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 09)*
