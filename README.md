# Go Mastery — A Complete Journey from Zero to Professional

> **Go Mastery** is a complete, self-contained, private tutorial for the Go programming language (a.k.a. Golang). It takes you from installing the toolchain to writing idiomatic, concurrent, production-grade Go — with source code, a full guide book, a cheat sheet, hands-on projects, exercises with solutions, conventions, and performance guidance.

---

## Table of Contents

### 📘 The Guide Book (`docs/`)
| # | Chapter | What you learn |
|---|---------|----------------|
| 01 | [Introduction](docs/01-introduction.md) | Why Go, history, design philosophy, installing the toolchain |
| 02 | [Hello, World](docs/02-hello-world.md) | `go run`, `go build`, package `main`, the module system |
| 03 | [Variables & Constants](docs/03-variables-constants.md) | `var`, `:=`, `const`, `iota`, scope, zero values |
| 04 | [Data Types](docs/04-data-types.md) | Integers, floats, strings, runes, booleans, conversions |
| 05 | [Operators](docs/05-operators.md) | Arithmetic, comparison, logical, bitwise, shifts |
| 06 | [Control Flow](docs/06-control-flow.md) | `if`, `for`, `switch`, `defer`, loops and labels |
| 07 | [Functions](docs/07-functions.md) | Parameters, multiple returns, variadics, closures, recursion |
| 08 | [Arrays & Slices](docs/08-arrays-slices.md) | Fixed arrays, dynamic slices, `append`, `copy`, capacity |
| 09 | [Maps](docs/09-maps.md) | Key/value stores, the comma-ok idiom, iteration order |
| 10 | [Structs](docs/10-structs.md) | Named fields, embedding, tags, value vs pointer semantics |
| 11 | [Methods](docs/11-methods.md) | Receivers, pointer vs value receivers, method sets |
| 12 | [Interfaces](docs/12-interfaces.md) | Implicit satisfaction, `error`, `io.Reader`, type switches |
| 13 | [Error Handling](docs/13-errors.md) | `error`, wrapping, `errors.Is/As`, custom errors, panics |
| 14 | [Pointers](docs/14-pointers.md) | Addresses, dereferencing, escape analysis, aliasing |
| 15 | [Packages & Modules](docs/15-packages-modules.md) | Exports, `go.mod`, visibility, project layout |
| 16 | [Goroutines](docs/16-goroutines.md) | Lightweight threads, the scheduler, common pitfalls |
| 17 | [Channels](docs/17-channels.md) | Buffered/unbuffered channels, direction, closing |
| 18 | [Concurrency Patterns](docs/18-concurrency-patterns.md) | `select`, `sync`, worker pools, pipelines, `context` |
| 19 | [Generics](docs/19-generics.md) | Type parameters, constraints, when (not) to use them |
| 20 | [File I/O](docs/20-file-io.md) | `os`, `bufio`, readers/writers, streaming |
| 21 | [JSON](docs/21-json.md) | Marshal/unmarshal, tags, streaming encoders |
| 22 | [Testing](docs/22-testing.md) | Unit tests, table-driven tests, benchmarks, fuzzing |
| 23 | [Web Server](docs/23-web-server.md) | `net/http`, handlers, middleware, REST |
| 24 | [Performance](docs/24-performance.md) | Profiling, allocations, escapes, benchmarks |
| 25 | [Best Practices](docs/25-best-practices.md) | Idiomatic Go, error discipline, API design |
| 26 | [Conventions & Idioms](docs/26-conventions-and-idioms.md) | `gofmt`, naming, comments, common idioms |
| 27 | [When to Use What](docs/27-when-to-use-what.md) | Decision tables: slice vs array, pointer vs value, etc. |
| 28 | [Next Steps](docs/28-next-steps.md) | Where to go after this book |

### 📄 Other resources
- [Cheat Sheet](cheat-sheet/CHEAT_SHEET.md) — one-page syntax and standard-library reference
- [Exercises & Questions](exercises/QUESTIONS.md) — 150+ problems across all levels
- [Exercise Solutions](exercises/SOLUTIONS.md) — detailed, explained answers
- [Examples](examples/) — runnable source code for every topic
- [Projects](projects/) — six progressively larger real programs

---

## How to use this repository

1. **Read a chapter** in `docs/`, then **run the matching example**:
   ```bash
   go run ./examples/08-arrays-slices
   ```
2. **Do the exercises** at the end of each chapter, then check `exercises/SOLUTIONS.md`.
3. **Build the projects** in `projects/` in order — each reinforces a new set of skills.
4. **Keep the cheat sheet open** while coding.

All examples live in a single Go module rooted at this directory, so from the repo root you can run or test any package:

```bash
go run ./examples/<folder>
go test ./...
gofmt -l .        # should list nothing
go vet ./...
```

> **Note:** Go is not required to *read* this tutorial, but you need it to *run* the code. Install Go 1.22+ from <https://go.dev/dl/>.

---

## Repository layout

```
go-mastery/
├── README.md                  <- you are here
├── go.mod                     <- single module for all runnable code
├── docs/                      <- the guide book (28 chapters)
├── cheat-sheet/               <- quick reference
├── examples/                  <- runnable snippets per topic
│   ├── 01-hello/
│   ├── 02-variables/
│   └── ...
├── projects/                  <- 6 hands-on projects
│   ├── 01-calculator/
│   ├── 02-todo-cli/
│   ├── 03-word-frequency/
│   ├── 04-concurrency-pipeline/
│   ├── 05-url-shortener/
│   └── 06-expense-tracker/
└── exercises/
    ├── QUESTIONS.md
    └── SOLUTIONS.md
```

---

## Learning path (recommended order)

```
Beginner      -> chapters 01-07  -> calculator project
Intermediate  -> chapters 08-15  -> todo-cli + word-frequency
Advanced      -> chapters 16-21  -> concurrency-pipeline + url-shortener
Professional  -> chapters 22-27  -> expense-tracker (full app with tests)
```

---

*Built for learning. Private repository. Ownership: `anubhav-1207`.*
