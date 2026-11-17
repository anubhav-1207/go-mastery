# Chapter 28 — Next Steps

You have completed the guide book. Here is how to consolidate and continue.

## Consolidate what you learned

1. Re-read the chapters you found hardest.
2. Rebuild each project from scratch without looking.
3. Re-do the exercises in `exercises/QUESTIONS.md` from memory.
4. Keep the cheat sheet handy for a week.

## Build more projects

Ideas that stretch different skills:

- **CLI**: a `grep`-like tool with regex and flags (`flag` package).
- **Web**: a blog with templates, authentication, and a SQLite database.
- **Concurrency**: a web crawler with rate limiting and deduplication.
- **Systems**: a TCP chat server with `net` and goroutines per connection.
- **Tooling**: a code generator using `go/ast` and `text/template`.
- **Cloud**: a REST service with `slog`, health checks, graceful shutdown, and Docker packaging.

## Essential standard library packages to explore

| Package | Purpose |
|---------|---------|
| `context` | cancellation, deadlines |
| `net/http` | HTTP client/server |
| `encoding/json` | JSON |
| `database/sql` | SQL databases |
| `log/slog` | structured logging |
| `os`, `io`, `bufio`, `path/filepath` | files and streams |
| `time` | time, timers, tickers |
| `sync`, `sync/atomic` | synchronization |
| `errors`, `fmt`, `strings`, `strconv` | fundamentals |
| `slices`, `maps`, `cmp` | generic collection helpers |
| `testing`, `testing/quick` | tests and property testing |
| `embed` | embed assets in binaries |
| `flag` | CLI flags |
| `regexp` | regular expressions |
| `go/ast`, `go/parser` | source analysis |

## Popular third-party libraries

- **Web routers**: `github.com/go-chi/chi`, `github.com/gin-gonic/gin`
- **Validation**: `github.com/go-playground/validator/v10`
- **Config**: `github.com/spf13/viper`, `github.com/caarlos0/env`
- **CLI**: `github.com/spf13/cobra`, `github.com/urfave/cli`
- **Testing**: `github.com/stretchr/testify`, `github.com/google/go-cmp`
- **Concurrency**: `golang.org/x/sync/errgroup`, `golang.org/x/sync/semaphore`
- **Database**: `github.com/jmoiron/sqlx`, `gorm.io/gorm`
- **Linting**: `golangci-lint`
- **Logging**: `go.uber.org/zap` (or stdlib `slog`)

## Read great code

- The Go standard library source (`$GOROOT/src`) — especially `net/http`, `io`, `sync`.
- Well-known projects: Docker, Kubernetes, Prometheus, Terraform, Caddy, Hugo.
- Effective Go: <https://go.dev/doc/effective_go>
- Go Code Review Comments: <https://go.dev/wiki/CodeReviewComments>

## Official resources

- Tour of Go: <https://go.dev/tour/>
- Go by Example: <https://gobyexample.com/>
- Go Blog: <https://go.dev/blog/>
- Go spec: <https://go.dev/ref/spec>
- pkg.go.dev: <https://pkg.go.dev/>
- Go Weekly newsletter: <https://golangweekly.com/>

## Practice platforms

- Exercism Go track
- LeetCode (Go)
- Codewars
- Advent of Code (in Go!)

## Keep growing

- Learn `go generate`, build tags, cgo basics.
- Understand the runtime: GC, scheduler, pprof deeply.
- Read "The Go Programming Language" (Donovan & Kernighan) and "Designing Data-Intensive Applications".
- Contribute to open source — even documentation fixes.

## Final words

Go rewards simplicity, clarity, and discipline. Write boring, obvious code that works, test it, measure it, and let composition carry you. Welcome to the tribe, my lord. 🐹

## Exercises

1. Pick one project idea and write a one-page design (types, packages, interfaces) before coding.
2. Read the source of `io.Copy` and explain its optimizations.
3. Complete one Exercism Go exercise per day for a week.
4. Contribute a small fix to an open-source Go project.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 28)*
