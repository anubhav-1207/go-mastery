# Chapter 15 — Packages & Modules

## Packages

A **package** is a directory of `.go` files compiled together. Every file starts with the same `package` declaration.

- Package name is usually the directory name (lowercase, short).
- `package main` builds an executable; any other name builds a library.
- Multiple files in one directory share the same package scope.

```go
// mathutil/mathutil.go
package mathutil

func Add(a, b int) int { return a + b }
```

Importing it:

```go
import "github.com/anubhav-1207/go-mastery/mathutil"

fmt.Println(mathutil.Add(2, 3))
```

## Exported vs unexported names

Visibility is determined by the **first letter**:

| First letter | Visibility |
|--------------|-----------|
| Uppercase (`Add`, `Person`, `MaxSize`) | **Exported** — visible outside the package |
| Lowercase (`add`, `person`, `maxSize`) | Package-private |

This applies to variables, constants, types, functions, methods, and struct fields.

```go
type User struct {
    Name   string // exported — readable/writable from other packages
    secret string // private
}
```

There is no `public`/`private`/`protected` keyword — capitalization is the whole rule.

## Importing

```go
import "fmt"

import (
    "fmt"
    "os"

    "github.com/gin-gonic/gin"   // third-party
)

import alias "github.com/long/path/pkg"   // alias
import _ "github.com/lib/pq"               // blank import: side effects only
import . "math"                            // dot import — AVOID (pollutes namespace)
```

- The compiler enforces that every import is used.
- Blank imports run the package's `init()` functions — used for registering database drivers.

## `init()` functions

A package may have `init()` functions executed automatically before `main`, in dependency order:

```go
func init() {
    // set up global state, register drivers, etc.
}
```

Use sparingly — explicit setup is easier to reason about and test.

## Modules

A **module** is a collection of packages versioned together, defined by a `go.mod` file.

```
module github.com/anubhav-1207/go-mastery

go 1.22

require github.com/some/lib v1.2.3
```

Key commands:

```bash
go mod init <path>     # create go.mod
go mod tidy            # sync dependencies: add missing, remove unused
go get pkg@version     # add/upgrade a dependency
go get pkg@latest      # latest version
go list -m all         # list the dependency graph
go mod verify          # verify downloaded modules
go mod download        # download modules
```

`go.sum` stores cryptographic hashes of dependency contents — commit it.

## Semantic versioning and pseudo-versions

Go modules follow semver: `vMAJOR.MINOR.PATCH`.

- `v1.2.3` — tagged release.
- `v0.x.y` — pre-1.0, no compatibility guarantees.
- `v1.2.4-0.20240101abc123` — pseudo-version for an untagged commit.

Modules with major version ≥ 2 include the major in the path: `github.com/foo/bar/v2`.

## Module resolution and `GOPROXY`

By default Go downloads modules from `proxy.golang.org`. Configure with:

```bash
go env -w GOPROXY=https://proxy.golang.org,direct
```

Set `GONOSUMDB`/`GOPRIVATE` for private modules:

```bash
go env -w GOPRIVATE=github.com/mycompany/*
```

## Project layout (common convention)

```
myapp/
├── cmd/
│   └── myapp/
│       └── main.go          # entry point(s)
├── internal/
│   ├── config/
│   ├── handler/
│   └── service/             # private to this module
├── pkg/
│   └── publiclib/           # reusable by external projects
├── api/                     # OpenAPI/protobuf definitions
├── test/                    # integration tests, fixtures
├── go.mod
├── go.sum
└── Makefile
```

- `cmd/` — one subdirectory per executable.
- `internal/` — packages that cannot be imported from outside the module (enforced by the compiler).
- `pkg/` — libraries safe for external use.

For small projects, a flat layout is fine. Do not over-engineer.

## Cycles

Import cycles are compile errors. Design packages as a DAG. If package A needs B and B needs A, extract shared types into a third package C.

## Runnable example

See `examples/14-packages/` which demonstrates a small multi-package module.

## Chapter summary

- Packages group files; capitalize names to export them.
- `package main` + `main()` builds an executable.
- Modules are versioned package trees defined by `go.mod` + `go.sum`.
- `go mod tidy` keeps dependencies clean.
- Use `cmd/`, `internal/`, `pkg/` for larger projects.
- Avoid import cycles and dot imports; use blank imports only for side effects.

## Exercises

1. Create a module with a `mathutil` package and a `main` that imports it.
2. Show that a lowercase function is not accessible from another package.
3. Add a third-party dependency and run `go mod tidy`.
4. Create an `internal` package and verify it can be imported locally.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 15)*
