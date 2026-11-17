# Chapter 02 — Hello, World

## The smallest Go program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Save it as `main.go`, then:

```bash
go run main.go
```

Expected output:
```
Hello, world!
```

## Dissecting it line by line

### `package main`
Every Go file begins with a **package declaration**. A package is a directory of `.go` files compiled together. The special name `main` means "this is an executable program" — it is not a library.

### `import "fmt"`
`fmt` (format) is a standard-library package for formatted I/O. `import` makes the package available in this file. Multiple imports use a block:

```go
import (
    "fmt"
    "os"
    "strings"
)
```

> **`gofmt` sorts imports for you** and groups standard library vs third-party with a blank line.

### `func main()`
`func` declares a function. `main` is special: when you run a `main` package, execution starts here. It takes no arguments and returns nothing.

### `fmt.Println(...)`
Calls the `Println` function from the `fmt` package: prints arguments followed by a newline.

## Compiling vs running

```bash
go run main.go          # compile to a temp binary, run it, discard it
go build -o hello .     # produce ./hello (or hello.exe on Windows)
./hello
```

Go produces a **statically linked native binary** by default — you can copy it to another machine with the same OS/architecture and it runs with no Go installation needed.

## Modules: the unit of dependency management

Since Go 1.16, modules are the default. A module is a directory tree with a `go.mod` file.

```bash
go mod init github.com/anubhav-1207/go-mastery
```

This creates:

```
module github.com/anubhav-1207/go-mastery

go 1.22
```

- `module` — the import path prefix for all packages in this tree.
- `go` — the minimum Go language version expected.

When you import third-party code, `go.mod` records the module and version, and `go.sum` records cryptographic checksums.

```bash
go get github.com/gorilla/mux@v1.8.1
go mod tidy    # remove unused, add missing
```

## `fmt` verbs you will use immediately

```go
name := "Gopher"
age := 7

fmt.Println(name, age)                    // Gopher 7
fmt.Printf("Name: %s, age: %d\n", name, age)
fmt.Printf("Value: %v, type: %T\n", name, name)
```

| Verb | Meaning |
|------|---------|
| `%v` | default format of any value |
| `%T` | type of the value |
| `%d` | decimal integer |
| `%f` | floating point |
| `%s` | string |
| `%q` | quoted string |
| `%t` | boolean |
| `%+v` | struct with field names |

## Comments

```go
// Single-line comment. Use full sentences, starting with the name of what you describe.

/*
   Multi-line comment (rarely used; prefer consecutive // lines).
*/
```

**Godoc convention:** a comment immediately above an exported identifier should begin with that identifier's name:

```go
// Greet returns a greeting for the given name.
func Greet(name string) string {
    return "Hello, " + name
}
```

Run `go doc` on your own packages to see these comments.

## Semicolons (why you don't write them)

Go inserts a semicolon automatically at the end of a line when the line ends with a token that can legally end a statement (identifier, literal, `return`, `)`, `]`, `}`). That is why braces must be on the same line as the `if`/`func`:

```go
if true {   // OK: the '{' must be on this line
}
```

Writing `if true\n{` would insert a semicolon after `true` and break the parse.

## Common beginner mistakes

| Mistake | Error / result |
|---------|----------------|
| `fmt.Printlf(...)` | `undefined: fmt.Printlf` |
| Unused import | compile error: `"fmt" imported and not used` |
| Unused local variable | compile error: `declared and not used` |
| Capitalized `main` function | no entry point |
| Missing `package main` | cannot build an executable |

Go treats unused imports and unused local variables as **errors**, not warnings. This keeps code clean.

## Runnable example

See `examples/01-hello/main.go`.

## Chapter summary

- A program is a `main` package with a `main()` function.
- Use `go run` during development and `go build` to produce a binary.
- A `go.mod` file defines a module and its dependencies.
- `gofmt` formats code; Go enforces it by convention.
- Unused imports/variables are compile errors.

## Exercises

1. Print your name, age, and favorite language on separate lines using `Println`.
2. Use `Printf` to print `I have 3 apples and 2 oranges` using `%d` verbs.
3. Build a binary named `greet` and run it.
4. Deliberately import an unused package and observe the compiler error.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 02)*
