# Chapter 27 — When to Use What

A decision guide for the choices that come up constantly in Go.

## Arrays vs slices

| Situation | Use |
|-----------|-----|
| Fixed number of elements known at compile time | Array `[N]T` |
| Need to pass to functions without copying | Slice (or `*[N]T`) |
| Dynamic size, `append`, most application code | Slice `[]T` |
| Matrix with fixed dimensions | `[M][N]T` or `[][]T` if ragged |
| Want value semantics / no sharing | Array |

**Default: use a slice.** Arrays are for special cases.

## Slice vs `[]byte` vs `string`

| Need | Use |
|------|-----|
| Immutable text | `string` |
| Mutable bytes / binary data | `[]byte` |
| Text you will modify | `[]byte` then convert back |
| Network I/O buffers | `[]byte` |
| Function parameter for read-only text | `string` |

## `map` vs slice of structs

| Need | Use |
|------|-----|
| Lookup by key, uniqueness | `map[K]V` |
| Ordered sequence, duplicates allowed | `[]T` |
| Need deterministic iteration | Slice (sort keys if using map) |
| Small set (< ~10 items) scanned linearly | Slice is simpler/faster |
| Set membership | `map[T]struct{}` |

## Pointer vs value receiver

| Situation | Receiver |
|-----------|----------|
| Method mutates receiver | `*T` |
| Struct is large | `*T` |
| Type contains mutex/file/buffer (non-copyable) | `*T` |
| Small, immutable struct | `T` |
| Mixed methods on a type | Use `*T` for **all** (consistency) |

## Pointer vs value parameter

| Situation | Pass |
|-----------|------|
| Need to modify the argument | pointer |
| Large struct | pointer |
| Small scalar | value |
| Want to allow `nil` meaning "absent" | pointer |
| Slice/map/channel (descriptor) | value (usually) |
| Need to reassign the descriptor itself | pointer or return it |

## `error` vs `panic`

| Situation | Use |
|-----------|-----|
| Expected failure (bad input, network error) | `error` |
| Caller can recover | `error` |
| Unrecoverable programmer bug / impossible state | `panic` |
| Initialization that must succeed | `panic` (or `log.Fatal` in main) |
| Library code | `error`, never `panic` |

## Goroutine vs synchronous call

| Situation | Use |
|-----------|------|
| Task is fast and CPU-light | synchronous |
| I/O bound, many concurrent operations | goroutine |
| Need parallelism for CPU work | goroutine (with `GOMAXPROCS`) |
| Need result immediately | synchronous |
| Fire-and-forget background task | goroutine + cancellation |

## Channel vs mutex

| Situation | Use |
|-----------|------|
| Passing data between goroutines / pipeline | channel |
| Protecting a shared value / invariant | mutex |
| Ownership is clear, data flows | channel |
| Many readers of shared state | `RWMutex` |
| Simple counter | `atomic` |

> "Channels are for orchestration; mutexes are for protection."

## Buffered vs unbuffered channel

| Situation | Use |
|-----------|------|
| Need rendezvous / strict synchronization | unbuffered |
| Smooth bursts, decouple producer/consumer | buffered |
| Don't know the right capacity | start unbuffered, measure |
| Intentional blocking (nil channel) | `nil` |

## `context` vs manual cancellation

| Situation | Use |
|-----------|------|
| Request-scoped, timeouts, deadlines, propagation | `context.Context` |
| Simple one-off stop signal | channel close |
| Library APIs doing I/O | accept `context.Context` |

## `any` vs generics vs interface

| Situation | Use |
|-----------|------|
| Truly heterogeneous values (`json.Unmarshal`) | `any` |
| Same algorithm for many types, type-safe | generics |
| Shared behavior across types | interface (methods) |
| Plug-in / dependency injection | interface |

## `init()` vs explicit setup

| Situation | Use |
|-----------|-----|
| Registering a driver / side effect on import | `init` (rare) |
| Anything testable or configurable | explicit function |

## `fmt.Println` vs `log` vs structured logging

| Situation | Use |
|-----------|-----|
| Quick debug / CLI output | `fmt.Println` |
| Server diagnostics with timestamps | `log` / `slog` |
| Production structured logs | `log/slog` (Go 1.21+) |

## Standard library vs third-party

| Situation | Use |
|-----------|-----|
| HTTP, JSON, file I/O, testing, concurrency | standard library |
| Router with complex needs, ORM, cloud SDK | vetted third-party |
| One-line utility you can write | write it yourself |

## `go run` vs `go build`

| Situation | Use |
|-----------|-----|
| Iterating quickly | `go run .` |
| Producing a deployable binary | `go build -o app .` |
| Installing a CLI tool | `go install` |

## Quick decision flowchart

```
Need to store things?
  ├─ lookup by key? -> map
  └─ sequence? -> slice (array only if fixed)

Need concurrency?
  ├─ passing data -> channel
  ├─ protecting state -> mutex/atomic
  └─ cancellation -> context

Need abstraction?
  ├─ behavior shared -> interface
  ├─ algorithm shared -> generics
  └─ runtime heterogeneity -> any

Something went wrong?
  ├─ recoverable -> error
  └─ unrecoverable -> panic
```

## Runnable example

See `examples/26-when-to-use-what/main.go` for side-by-side comparisons.

## Exercises

1. For each of ten scenarios, choose slice vs map and justify it.
2. Convert a value-receiver method to a pointer receiver and explain the method-set change.
3. Decide channel vs mutex for a counter, a job queue, and a config cache.
4. Refactor a function using `any` into a generic equivalent.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 27)*
