# Chapter 05 — Operators

## Arithmetic operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `+` | add | `3 + 2 == 5` |
| `-` | subtract | `3 - 2 == 1` |
| `*` | multiply | `3 * 2 == 6` |
| `/` | divide | `7 / 2 == 3` (integer division truncates) |
| `%` | remainder | `7 % 2 == 1` |

```go
fmt.Println(7 / 2)    // 3   (both operands int)
fmt.Println(7 / 2.0)  // 3.5 (one operand is float)
fmt.Println(-7 / 2)   // -3  (truncates toward zero)
fmt.Println(-7 % 2)   // -1  (sign follows dividend)
```

Increment/decrement are **statements**, not expressions:

```go
i := 0
i++    // OK
i--    // OK
// j := i++  // ERROR: ++ is not an expression
```

`+=`, `-=`, `*=`, `/=`, `%=` are assignment operators:

```go
x := 10
x += 5   // x = 15
x *= 2   // x = 30
```

## Comparison operators

`==`, `!=`, `<`, `<=`, `>`, `>=`

They work on numbers, strings (lexicographic byte order), booleans, pointers, and comparable structs/arrays. They return a `bool`.

```go
fmt.Println("apple" < "banana")  // true
fmt.Println(3 == 3.0)            // true? NO — compile error: mismatched types int and float64
```

> Comparison requires identical types. Convert explicitly: `float64(3) == 3.0`.

**Slices, maps, and functions can only be compared to `nil`** — use `reflect.DeepEqual` or `slices.Equal` for deep comparison.

## Logical operators

`&&` (and), `||` (or), `!` (not). Both `&&` and `||` **short-circuit**:

```go
if x != 0 && 10/x > 2 {  // safe: 10/x is only evaluated if x != 0
    // ...
}
```

## Bitwise operators

| Operator | Meaning |
|----------|---------|
| `&` | AND |
| `\|` | OR |
| `^` | XOR |
| `&^` | AND NOT (clear bits) |
| `<<` | left shift |
| `>>` | right shift |

```go
a := 0b1100  // 12
b := 0b1010  // 10

fmt.Printf("%04b\n", a&b)  // 1000  (8)
fmt.Printf("%04b\n", a|b)  // 1110  (14)
fmt.Printf("%04b\n", a^b)  // 0110  (6)
fmt.Printf("%04b\n", a&^b) // 0100  (4) — clear bits set in b

fmt.Println(1 << 10)       // 1024
fmt.Println(1024 >> 3)     // 128
```

Unary `^` is bitwise complement: `^0` is `-1` (all bits set in two's complement).

Bitwise operators are common with flags and low-level protocols:

```go
const (
    Read  = 1 << iota  // 1
    Write                // 2
    Execute              // 4
)

perm := Read | Write    // 3
fmt.Println(perm&Read != 0)    // true — has Read
fmt.Println(perm&Execute != 0) // false — no Execute
```

## Operator precedence

From highest to lowest:

```
5  *  /  %  <<  >>  &  &^
4  +  -  |  ^
3  ==  !=  <  <=  >  >=
2  &&
1  ||
```

Use parentheses when in doubt — they are free and improve readability:

```go
if (a && b) || c { ... }
```

## String concatenation

`+` concatenates strings; `+=` appends. **In a loop this is inefficient** because strings are immutable and each `+` allocates a new string. Use `strings.Builder`:

```go
var sb strings.Builder
for i := 0; i < 1000; i++ {
    sb.WriteString("item ")
    sb.WriteString(strconv.Itoa(i))
    sb.WriteByte('\n')
}
result := sb.String()
```

## Runnable example

See `examples/04-operators/main.go`.

## Chapter summary

- `/` on integers truncates; `%` gives the remainder.
- `++`/`--` are statements, not expressions.
- Comparison requires identical, comparable types.
- `&&`/`||` short-circuit — use that for safety checks.
- Use bit shifts and masks for flags and power-of-two math.
- Build strings with `strings.Builder` inside loops.

## Exercises

1. Evaluate `^uint8(0)` and explain the result.
2. Use a bitmask to store permissions and test for a flag.
3. Show that `"Z" < "a"` using byte comparison.
4. Rewrite a loop that concatenates strings with `+` to use `strings.Builder`.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 05)*
