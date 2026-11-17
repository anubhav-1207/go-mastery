# Go Mastery — Questions & Exercises

**152 problems**, organized by chapter. Each chapter has conceptual questions (C) and coding tasks (P).
Detailed answers are in [SOLUTIONS.md](SOLUTIONS.md). Try every problem before peeking.

---

## Chapter 01 — Introduction

**C1.1** Why was Go created? Name three problems it was designed to solve at Google.
**C1.2** Explain the Go philosophy of "composition over inheritance." How do structs and interfaces support it?
**C1.3** Name four domains where Go excels and two where it is a weaker choice. Why?
**P1.4** Install Go, verify with `go version`, and print `go env GOROOT` and `go env GOPATH`.
**P1.5** Use `go doc fmt.Println` to read its documentation. Then write what the `...any` parameter means.

## Chapter 02 — Hello, World

**C2.1** What does `package main` signify? What happens if a library package declares `func main()`?
**C2.2** Why does Go treat unused imports and unused local variables as errors instead of warnings?
**C2.3** Explain why the opening brace must be on the same line as `func main()`.
**P2.4** Write a program that prints your name, age, and favorite language using `Printf` with `%s`, `%d`.
**P2.5** Build a binary named `greet` and run it. Then deliberately import an unused package and record the error.

## Chapter 03 — Variables & Constants

**C3.1** List the zero values for `int`, `float64`, `bool`, `string`, pointer, slice, map, and struct.
**C3.2** Explain the rule for redeclaring variables with `:=` in a multi-variable assignment.
**C3.3** What is `iota`? What value does `KB` get in `const ( _ = iota; KB = 1 << (10*iota) )`?
**P3.4** Declare variables using both `var` and `:=`, print their zero values, then demonstrate shadowing.
**P3.5** Create constants for the four seasons starting at 1 using `iota`, and show a constant assigned to both `int` and `float64`.

## Chapter 04 — Data Types

**C4.1** Why should you default to `int` and `float64`? When would you choose `int64` or `uint8`?
**C4.2** Explain the difference between `byte`, `rune`, and `string`. Why is `len("世界")` not 2?
**C4.3** What is the difference between `type MyInt int` and `type MyInt = int`?
**P4.4** Print the byte length and rune length of `"Go 语言"`, then iterate with `range` printing index, rune, and code point.
**P4.5** Reverse a string containing multibyte characters using `[]rune`. Convert `int(300)` to `uint8` and explain the result.

## Chapter 05 — Operators

**C5.1** Explain integer division and the sign of the remainder for `-7 % 2`.
**C5.2** Why are `++` and `--` statements rather than expressions? Give an example that would fail if they were expressions.
**C5.3** What does short-circuit evaluation buy you? Give a safe-check example.
**P5.4** Use bit shifts to create flags `Read`, `Write`, `Execute`; build a permission value and test membership.
**P5.5** Rewrite a loop that concatenates strings with `+` to use `strings.Builder`, and explain why.

## Chapter 06 — Control Flow

**C6.1** Go has only `for`. Show the four forms it can take.
**C6.2** Why is map iteration order unspecified? What must you do to print map keys deterministically?
**C6.3** Describe three properties of `defer` (order, argument evaluation, panic behavior).
**P6.4** Write a labeled `break` that exits two nested loops when `i*j > 10`.
**P6.5** Write a `switch` (expressionless form) that categorizes a number as negative, zero, small (1–10), or large (>10).

## Chapter 07 — Functions

**C7.1** What are the advantages of multiple return values over exceptions?
**C7.2** What is a closure? What does the following print and why?
```go
func counter() func() int { n := 0; return func() int { n++; return n } }
c := counter(); fmt.Println(c(), c(), c())
```
**C7.3** When are naked returns acceptable? When should you avoid them?
**P7.4** Write a variadic `average(nums ...float64) float64` returning 0 for no arguments.
**P7.5** Write `minMax(nums []int) (min, max int)` and a higher-order `apply(fn func(int) int, xs []int) []int`.

## Chapter 08 — Arrays & Slices

**C8.1** Explain the difference between an array and a slice in terms of copying and length.
**C8.2** What are `len` and `cap`? When does `append` reallocate? Why must you assign its result?
**C8.3** What is the difference between a `nil` slice and an empty slice? Where does it matter?
**P8.4** Write `removeAt(s []int, i int) []int` preserving order. Then demonstrate two slices sharing a backing array.
**P8.5** Use `slices.SortFunc` to sort a `[]Person` by `Age`, and use `slices.Clone` to detach a slice.

## Chapter 09 — Maps

**C9.1** Why can slices not be map keys? Which types can be keys?
**C9.2** Explain the comma-ok idiom. Why is `v := m[k]` insufficient to test existence?
**C9.3** Why is map iteration randomized? How do you iterate deterministically?
**P9.4** Count word frequencies in `[]string` and implement a set with `map[string]struct{}`.
**P9.5** Group `[]string` words by first letter into `map[string][]string`, then clear the map.

## Chapter 10 — Structs

**C10.1** What are struct tags? Which package uses them heavily and how?
**C10.2** Explain embedding and field/method promotion. How is it different from inheritance?
**C10.3** Why can `struct{ Data []int }` not be compared with `==`?
**P10.4** Define `Rectangle` with `Area()` as a function (not method), then show value vs pointer assignment behavior.
**P10.5** Add JSON tags to a struct so one field is omitted when empty and another is never serialized.

## Chapter 11 — Methods

**C11.1** Compare value and pointer receivers. When is each appropriate?
**C11.2** Why should a type with one pointer-receiver method use pointer receivers for all methods?
**C11.3** What is a method set? How do the method sets of `T` and `*T` differ?
**P11.4** Implement `BankAccount` with `Deposit`/`Withdraw` (pointer receivers) and `Balance()` (value receiver).
**P11.5** Implement `String() string` for a custom type and show how `fmt` uses it. Demonstrate a nil receiver.

## Chapter 12 — Interfaces

**C12.1** What does "implicitly satisfied" mean? Why is this useful?
**C12.2** Explain the typed-nil-in-interface pitfall with an example.
**C12.3** What does "accept interfaces, return structs" mean and why is it idiomatic?
**P12.4** Define a `Shape` interface with `Area()`, implement it for `Circle` and `Square`, and compute total area.
**P12.5** Write a type switch for `int`, `float64`, `string`, and `nil`. Add a compile-time interface check.

## Chapter 13 — Errors

**C13.1** Why does Go return errors as values instead of using exceptions?
**C13.2** Explain `%w` vs `%v` in `fmt.Errorf` and the roles of `errors.Is` and `errors.As`.
**C13.3** When is `panic` acceptable? Why should library code avoid it?
**P13.4** Write a function that wraps an OS error with the file path and check it with `errors.Is(err, os.ErrNotExist)`.
**P13.5** Define a custom error type with a field, wrap it, and extract it with `errors.As`. Write a `recover` wrapper converting a panic to an error.

## Chapter 14 — Pointers

**C14.1** Explain `&x`, `*p`, and `*int`. What is the zero value of a pointer?
**C14.2** What is escape analysis? Why does a returned pointer to a local variable still work?
**C14.3** When do you need a pointer to a slice? Give an example where a value receiver suffices.
**P14.1** Write `swap(a, b *int)` and demonstrate pointer aliasing.
**P14.2** Write `grow(s *[]int)` that appends so the caller sees the new slice; compare with returning the slice.

## Chapter 15 — Packages & Modules

**C15.1** How does Go determine whether an identifier is exported?
**C15.2** What are the roles of `go.mod` and `go.sum`? What does `go mod tidy` do?
**C15.3** What is the purpose of `internal/` and `cmd/` directories? What is an import cycle?
**P15.4** Create a module with a `mathutil` package and a `main` that imports it. Show that a lowercase function is inaccessible.
**P15.5** Write a package with an `init()` function and a blank import that triggers a side effect.

## Chapter 16 — Goroutines

**C16.1** What is a goroutine? Why can you start far more goroutines than OS threads?
**C16.2** Describe the GMP scheduler model briefly.
**C16.3** List four common goroutine pitfalls and their fixes.
**P16.4** Launch 10 goroutines that print their ID, waiting with a `WaitGroup`. Ensure `Add` is called before `go`.
**P16.5** Create a data race with a shared counter, then fix it with a mutex. Run `go run -race`.

## Chapter 17 — Channels

**C17.1** Compare unbuffered and buffered channels. When is each useful?
**C17.2** Who should close a channel? What happens if you send on a closed channel or receive from a nil channel?
**C17.3** What are fan-out and fan-in? How do channels provide backpressure?
**P17.4** Write a Fibonacci generator emitting on a channel and a fan-out that distributes jobs to 3 workers.
**P17.5** Write a `merge(cs ...<-chan int) <-chan int` fan-in function and test it with two channels.

## Chapter 18 — Concurrency Patterns

**C18.1** What does `select` do? What is the role of `default` and an empty `select {}`?
**C18.2** Why should `ctx` be the first parameter? Why must you always call `cancel()`?
**C18.3** Channel or mutex? Give three scenarios each.
**P18.4** Implement a worker pool with 4 workers and 20 jobs, closing channels correctly.
**P18.5** Use `context.WithTimeout` to cancel a pipeline stage, and use `sync.Once` for one-time setup.

## Chapter 19 — Generics

**C19.1** What problem do generics solve? Show a pre-generics and a generic version of `Max`.
**C19.2** What do `~int` and `|` mean in a constraint? What is `comparable`?
**C19.3** When should you avoid generics?
**P19.1** Write `Filter[T any]` and `Map[T, U any]`.
**P19.2** Implement `Set[T comparable]` with Add, Contains, Remove.
**P19.3** Write a generic `Stack[T]` with pointer-receiver methods.
**P19.4** Explain why a method cannot have its own type parameters.
**P19.5** Use `cmp.Ordered` to write `Min[T](a, b T) T` and test with ints and strings.

## Chapter 20 — File I/O

**C20.1** What are `io.Reader` and `io.Writer`? Why are they powerful?
**C20.2** Why must you handle `n > 0` before checking `io.EOF`?
**C20.3** Why is `filepath.Join` better than string concatenation for paths?
**P20.4** Count lines in a file with `bufio.Scanner` and copy a binary file with `io.Copy`.
**P20.5** Write 10,000 lines with and without `bufio.Writer`; compare timing. Use `os.Stat` to print size and mod time.

## Chapter 21 — JSON

**C21.1** What do the tags `json:"name"`, `omitempty`, and `json:"-"` do?
**C21.2** Why must you pass a pointer to `json.Unmarshal`? What type do JSON numbers decode to by default?
**C21.3** When would you use `json.Encoder`/`Decoder` instead of `Marshal`/`Unmarshal`?
**P21.1** Marshal a struct with `omitempty` and inspect the output.
**P21.2** Unmarshal with `DisallowUnknownFields` and show the error.
**P21.3** Implement `MarshalJSON`/`UnmarshalJSON` for a custom `Money` type.
**P21.4** Write a streaming decoder for newline-delimited JSON.
**P21.5** Explain why `map[string]any` is risky and show the safe alternative.

## Chapter 22 — Testing

**C22.1** Why are table-driven tests the Go idiom? What does `t.Run` give you?
**C22.2** What do `t.Helper`, `t.TempDir`, and `t.Cleanup` do?
**C22.3** What do `ns/op`, `B/op`, and `allocs/op` mean? Why use `-benchmem`?
**P22.1** Write a table-driven test for `Divide` including an error case.
**P22.2** Add a benchmark comparing `+` and `strings.Builder`.
**P22.3** Write an `Example` function and a fuzz test for `Reverse`.
**P22.4** Write a test using `httptest` for a handler returning 404.
**P22.5** Run `go test -race -cover ./...` and interpret the output.

## Chapter 23 — Web Server

**C23.1** What is `http.Handler`? How does `http.HandlerFunc` help?
**C23.2** Why must headers be set before `WriteHeader`/`Write`?
**C23.3** What is middleware? Give four examples of useful middleware.
**P23.1** Build a server with `GET /hello?name=...` that greets the name.
**P23.2** Add middleware that sets an `X-Request-Id` header.
**P23.3** Implement `POST /echo` returning the JSON request body.
**P23.4** Write an `httptest` test for a handler returning 404.
**P23.5** Add graceful shutdown on SIGINT and set `ReadHeaderTimeout`.

## Chapter 24 — Performance

**C24.1** What is the correct optimization workflow?
**C24.2** What do escape analysis and `pprof` tell you?
**C24.3** List five common allocation-reduction techniques.
**P24.1** Benchmark `append` with and without preallocation; explain the difference.
**P24.2** Profile a slow function with `pprof` and identify the hotspot.
**P24.3** Use `testing.AllocsPerRun` to verify a function is allocation-free.
**P24.4** Compare `json.Marshal` with a hand-written encoder for a hot struct.
**P24.5** Explain why reducing allocation is usually better than tuning `GOGC`.

## Chapter 25 — Best Practices

**C25.1** Why should you wrap errors with context but log only once?
**C25.2** Explain "accept interfaces, return structs" with an example.
**C25.3** Why should every goroutine have a termination path?
**P25.1** Refactor a deeply nested function into guard clauses.
**P25.2** Convert code that ignores errors into wrapped-error code.
**P25.3** Design a service with an interface dependency and a fake for tests.
**P25.4** Run the pre-commit checklist on one of the projects.

## Chapter 26 — Conventions & Idioms

**C26.1** List five naming conventions with examples.
**C26.2** What is the blank-identifier compile-time interface check?
**C26.3** What are functional options and when are they useful?
**P26.1** Rewrite `HTTPHttpServer` idiomatically.
**P26.2** Convert a constructor with three boolean flags to functional options.
**P26.3** Add compile-time interface checks for two implementations.
**P26.4** Replace a `switch` on a `type` string with an interface-based design.

## Chapter 27 — When to Use What

**C27.1** Slice vs array vs map: give a scenario for each.
**C27.2** Channel vs mutex: give three scenarios each.
**C27.3** `any` vs generics vs interface: give a scenario for each.
**P27.1** For ten scenarios, choose a data structure and justify it.
**P27.2** Convert a value-receiver method to pointer receiver and explain the method-set change.
**P27.3** Decide channel vs mutex for a counter, a job queue, and a config cache.
**P27.4** Refactor an `any`-based function into a generic equivalent.

## Chapter 28 — Next Steps

**C28.1** Name five standard-library packages you have not used yet and what they do.
**C28.2** What is `slog`? Why prefer it over `fmt.Println` in servers?
**P28.1** Design (on paper) a blog API: types, packages, interfaces.
**P28.2** Read the source of `io.Copy` and summarize its optimizations.
**P28.3** Complete one Exercism Go exercise and describe what it taught you.

---

## Bonus challenges (B1–B12)

**B1** Implement a thread-safe LRU cache with `container/list` and `sync.Mutex`.
**B2** Write a generic `Ring[T]` buffer.
**B3** Build a concurrent web crawler with deduplication and a depth limit.
**B4** Implement a rate limiter using a token bucket.
**B5** Write a `mapreduce`-style word count over multiple files using goroutines.
**B6** Create a middleware chain type: `type Middleware func(http.Handler) http.Handler` with `Chain(mws...)`.
**B7** Implement graceful shutdown with `os/signal` and a 5-second deadline.
**B8** Write a custom `error` type that carries a stack trace (capture with `runtime.Callers`).
**B9** Build a CLI with `cobra`-style subcommands using only the standard library.
**B10** Use `embed` to serve static files from the binary.
**B11** Write a fuzz test that finds a bug in a hand-written `Atoi` clone.
**B12** Profile an intentionally slow program and produce a before/after benchmark report.

---

*Solutions: [SOLUTIONS.md](SOLUTIONS.md)*
