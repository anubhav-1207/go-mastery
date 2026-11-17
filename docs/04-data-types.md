# Chapter 04 — Data Types

Go is **statically typed**: every variable has a type known at compile time, and types are not automatically converted.

## Basic types

### Integers

| Type | Range |
|------|-------|
| `int` | platform-dependent: 32-bit or 64-bit (use this by default) |
| `int8` | -128 to 127 |
| `int16` | -32768 to 32767 |
| `int32` | -2³¹ to 2³¹-1 |
| `int64` | -2⁶³ to 2⁶³-1 |
| `uint` | platform-dependent unsigned |
| `uint8` (alias `byte`) | 0 to 255 |
| `uint16`, `uint32`, `uint64` | corresponding unsigned ranges |
| `uintptr` | integer large enough to hold a pointer (rarely used directly) |

```go
var a int = 42
var b uint8 = 255
var c byte = 'A'   // byte is an alias for uint8
```

**Rule of thumb:** use `int` unless you need a specific size (network protocols, binary formats, memory-critical structs).

### Floating point

| Type | Description |
|------|-------------|
| `float32` | 32-bit (~7 decimal digits) |
| `float64` | 64-bit (~15 digits) — **default** |

```go
var pi float64 = 3.14159
var small float32 = 0.5
```

### Complex numbers

```go
var c1 complex64 = 1 + 2i
var c2 complex128 = complex(3, 4)  // 3 + 4i
real(c2)  // 3
imag(c2)  // 4
```

### Boolean

```go
var isReady bool = true
```

Booleans do **not** convert to integers: `if 1 {}` is a compile error.

### Strings

Strings are **immutable sequences of bytes** (UTF-8 encoded by convention).

```go
s := "Hello, 世界"
len(s)        // 13 bytes (each CJK char = 3 bytes)
s[0]          // 'H' (a byte)
s + "!"       // concatenation
```

Useful `strings` package functions:

```go
strings.ToUpper(s)
strings.Contains(s, "世界")
strings.Split("a,b,c", ",")   // []string{"a","b","c"}
strings.Join([]string{"a","b"}, "-")
strings.TrimSpace("  hi  ")
strings.ReplaceAll(s, "old", "new")
strings.HasPrefix(s, "He")
```

### Runes

A `rune` is an alias for `int32` and represents a **Unicode code point**. Use runes when you need characters, not bytes:

```go
s := "Hello, 世界"
for i, r := range s {
    fmt.Printf("%d: %c (U+%04X)\n", i, r, r)
}
```

`range` over a string yields byte index + rune. Convert with:

```go
runes := []rune(s)   // []rune{'H','e','l','l','o',',',' ','世','界'}
len(runes)           // 9 characters (vs len(s) == 13 bytes)
```

## Type conversions

Go requires **explicit** conversion:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// No implicit conversion:
// var x float64 = i  // compile error
```

Converting between numeric types may truncate or wrap:

```go
int(3.99)    // 3  (truncates toward zero)
uint8(300)   // 44 (wraps around)
```

Converting `string` ↔ `[]byte` / `[]rune`:

```go
b := []byte("abc")  // copy of bytes
s := string(b)
r := []rune("世界")
```

## Composite (aggregate) types

| Type | Description | Chapter |
|------|-------------|---------|
| Array | fixed-length sequence | 08 |
| Slice | dynamic-length sequence | 08 |
| Map | hash table | 09 |
| Struct | collection of named fields | 10 |
| Pointer | memory address | 14 |
| Channel | concurrency conduit | 17 |
| Interface | set of methods | 12 |
| Function | function value | 07 |

## Type aliases vs defined types

```go
type MyInt int        // a NEW type; methods can be added, conversions required
type Byte = uint8     // alias: Byte IS uint8, identical type
```

```go
var a MyInt = 5
var b int = int(a)   // conversion needed
var c Byte = 200     // no conversion; Byte == uint8
```

## `type` declarations for clarity

```go
type UserID int64
type Email string

func Send(to Email) { /* ... */ }

// Send(Email("a@b.com")) — type safety prevents mixing IDs and emails
```

Named types make APIs self-documenting and prevent accidental mixing of values that share an underlying representation.

## Formatting values

```go
fmt.Printf("%d %b %x\n", 255, 255, 255)  // 255 11111111 ff
fmt.Printf("%.2f\n", 3.14159)            // 3.14
fmt.Printf("%q\n", "hi")                 // "hi"
fmt.Printf("%c\n", 'A')                  // A
```

## Runnable example

See `examples/03-data-types/main.go`.

## Chapter summary

- Prefer `int` and `float64` by default; choose fixed-size types for binary/network data.
- `byte` = `uint8`, `rune` = `int32` (a Unicode code point).
- Strings are immutable byte sequences; `range` yields runes.
- All conversions must be explicit; truncation is silent, so be careful.
- `type X int` creates a distinct type; `type X = int` creates an alias.

## Exercises

1. Print the byte length and rune length of `"Go 语言"`.
2. Convert `int(300)` to `uint8` and explain the result.
3. Create a named type `Age int` and write a function that accepts only `Age`.
4. Use `[]rune` to reverse a string containing multibyte characters.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 04)*
