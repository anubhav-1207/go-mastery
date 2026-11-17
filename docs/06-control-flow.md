# Chapter 06 — Control Flow

## `if`

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```

Unlike C, Go requires the condition to be a `bool` — `if score {}` is a compile error.

Go allows a **short statement** before the condition, separated by `;`. Variables declared there are scoped to the `if`/`else` chain:

```go
if err := doSomething(); err != nil {
    return err
} else {
    fmt.Println("ok")
}
```

This idiom is extremely common — it keeps `err` from leaking into the surrounding scope.

## `for` — the only loop keyword

Go has no `while` or `do-while`; `for` covers all forms.

### Classic three-clause form

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### Condition-only form (like `while`)

```go
i := 0
for i < 10 {
    i++
}
```

### Infinite loop

```go
for {
    // break to exit
}
```

### `range` form

Iterates over arrays, slices, strings, maps, and channels:

```go
nums := []int{10, 20, 30}
for i, v := range nums {
    fmt.Println(i, v)
}

for _, v := range nums {   // ignore the index
    fmt.Println(v)
}

for i := range nums {      // only the index
    fmt.Println(i)
}
```

For maps, `range` yields key → value in **random order** (intentionally unspecified):

```go
for key, value := range map[string]int{"a": 1, "b": 2} {
    fmt.Println(key, value)
}
```

For strings, the first value is the byte index and the second is the rune:

```go
for index, r := range "Go" {
    fmt.Println(index, r)
}
```

### `break` and `continue`

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue   // skip this iteration
    }
    if i == 8 {
        break      // exit the loop
    }
}
```

With **labels** you can break out of nested loops:

```go
Outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i*j == 4 {
            break Outer
        }
    }
}
```

## `switch`

Go's `switch` is more flexible than C's:

- Conditions need not be constants or integers.
- Cases do **not** fall through by default.
- You can switch on **any comparable type**.
- You can omit the expression entirely (a `switch true`).

```go
day := "Tuesday"
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Midweek")
}
```

### Expressionless switch

```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

### Short statement in switch

```go
switch v := getValue(); {
case v < 0:
    fmt.Println("negative")
case v == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive", v)
}
```

### `fallthrough`

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough   // deliberately continue to next case
case 2:
    fmt.Println("two")
}
```

If `x == 1`, both "one" and "two" print. Use `fallthrough` sparingly.

### Type switch

Used with interfaces to inspect dynamic type (Chapter 12):

```go
switch v := value.(type) {
case int:
    fmt.Println("int:", v*2)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
```

## `defer`

`defer` schedules a function call to run when the **enclosing function returns** — used for cleanup.

```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()    // guaranteed to run on return

    // ... use f ...
    return nil
}
```

Key behaviors:

1. **LIFO order** — last deferred runs first:
   ```go
   defer fmt.Println("first")
   defer fmt.Println("second")
   defer fmt.Println("third")
   // Output: third, second, first
   ```

2. **Arguments evaluated immediately**:
   ```go
   i := 1
   defer fmt.Println(i)  // prints 1
   i = 2
   ```

3. **Runs on panic** as the stack unwinds — perfect for releasing resources.

4. Defer in a loop can accumulate; extract into a helper if you open many files.

**Common pitfall — ignoring Close errors:**

```go
defer f.Close()   // return value discarded
```

If closing can fail meaningfully (e.g., buffered writes), handle it explicitly or name the return:

```go
func write() (err error) {
    f, _ := os.Create("out.txt")
    defer func() { err = errors.Join(err, f.Close()) }()
    _, err = f.WriteString("data")
    return err
}
```

## `goto` (rarely used)

Go has `goto`, but it cannot jump over variable declarations and is discouraged. Labels are mainly used with `break`/`continue`.

## Runnable example

See `examples/05-control-flow/main.go`.

## Chapter summary

- `if` supports an init statement; the condition must be a `bool`.
- `for` is the only loop keyword: classic, while-style, infinite, and `range`.
- `range` on maps is unordered; on strings it yields byte index + rune.
- `switch` does not fall through by default; `switch {}` is a clean alternative to `if/else` chains.
- `defer` runs LIFO on return — use it for cleanup.

## Exercises

1. Print even numbers 1–20 using `continue`.
2. Use a labeled `break` to exit two nested loops.
3. Write a `switch` that categorizes a number as negative/zero/positive/small/large.
4. Show with `defer` that execution order is LIFO.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 06)*
