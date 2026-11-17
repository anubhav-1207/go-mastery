# Chapter 01 — Introduction to Go

## What is Go?

Go (often called **Golang**) is an open-source programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and released in 2009. It was designed to solve problems Google faced with large-scale software: slow builds, complicated dependencies, and the difficulty of writing concurrent programs.

Go's motto could be summarized as: **simple, fast, and productive**.

## Design philosophy

Go deliberately removes features that add complexity without enough benefit:

| Go has | Go does **not** have |
|--------|----------------------|
| Structs (composition) | Classes / inheritance |
| Interfaces (implicit) | Inheritance hierarchies |
| Goroutines + channels | Heavyweight OS threads as the default |
| Garbage collection | Manual `free()` / `delete()` |
| Multiple return values | Exceptions (`try`/`catch`) |
| A single formatting style (`gofmt`) | Endless style debates |
| Generics (since 1.18) | Macro metaprogramming |

**Core principles:**
1. **Simplicity** — one obvious way to do things.
2. **Composition over inheritance** — embed small pieces to build big ones.
3. **Concurrency as a first-class citizen** — "Don't communicate by sharing memory; share memory by communicating."
4. **Fast compilation** — a large project builds in seconds.
5. **Batteries included** — a rich standard library (`net/http`, `encoding/json`, `testing`, ...).

## What is Go good at?

- **Backend APIs and microservices** — the `net/http` package is production-grade.
- **CLI tools** — single static binaries with no runtime dependencies.
- **Cloud infrastructure** — Docker, Kubernetes, Terraform, Prometheus are written in Go.
- **Concurrent data processing** — pipelines, fan-out/fan-in, worker pools.
- **Developer tooling** — `gofmt`, `go vet`, race detector are built in.

Go is *less* ideal for GUI desktop apps, real-time systems with hard latency requirements (GC pauses), and numerically heavy GPU computing (though it is improving).

## Installing Go

1. Download the installer for your OS from **<https://go.dev/dl/>**.
2. Install it, then verify in a terminal:

```bash
go version
# go version go1.27.4 linux/arm64
```

3. Set up your workspace (Go 1.16+ uses modules, so there is no required `GOPATH` layout):

```bash
mkdir -p ~/go-projects && cd ~/go-projects
go mod init example.com/hello
```

### Key environment variables

| Variable | Meaning | Typical value |
|----------|---------|---------------|
| `GOROOT` | Where Go itself is installed | set by installer |
| `GOPATH` | Workspace for cached modules/binaries | `~/go` |
| `GOMODCACHE` | Downloaded module cache | `~/go/pkg/mod` |
| `GOFLAGS` | Default flags for all commands | `-mod=mod` |

Check them with:

```bash
go env GOROOT GOPATH GOMODCACHE
```

## The Go toolchain (commands you will use daily)

```bash
go run main.go        # compile and run immediately
go build              # compile to a binary in the current directory
go build -o app .     # choose the output name
go install            # compile and place binary in $GOPATH/bin
go test ./...         # run all tests in the module
go vet ./...          # static analysis: catch suspicious code
gofmt -w .            # format all source files
go mod tidy           # clean up go.mod / go.sum
go get pkg@version    # add or upgrade a dependency
go doc fmt.Println    # read documentation from the terminal
```

## Your first mental model

A Go program is a collection of **packages**. One package's `main` function is the entry point. The compiler links everything into a single native binary.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

That is the whole program. No semicolons, no boilerplate class, no build system file besides `go.mod`.

## Chapter summary

- Go was built at Google for large-scale, concurrent, maintainable software.
- It favors simplicity, composition, fast builds, and a powerful standard library.
- Install Go 1.22+, verify with `go version`.
- `go run`, `go build`, `go test`, `gofmt`, and `go vet` are your core tools.

## Exercises

1. Install Go and print its version.
2. Explain in your own words why Go has no classes.
3. Run `go env` and find your `GOPATH` and `GOROOT`.
4. Use `go doc` to read the documentation for the `fmt` package's `Printf` function.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 01)*
