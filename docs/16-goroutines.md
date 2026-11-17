# Chapter 16 — Goroutines

A **goroutine** is a lightweight thread managed by the Go runtime. It is the foundation of Go's concurrency model.

## Starting a goroutine

Prefix any function call with `go`:

```go
func say(s string) {
    for i := 0; i < 3; i++ {
        fmt.Println(s, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

`go say("world")` runs concurrently; `main` continues immediately. Since the program exits when `main` returns, the goroutine may not finish — we will use channels (Chapter 17) or `sync.WaitGroup` (Chapter 18) to wait.

## Goroutines vs OS threads

| Property | Goroutine | OS thread |
|----------|-----------|-----------|
| Initial stack | ~2 KB, grows dynamically | ~1–2 MB fixed |
| Creation cost | very low | high |
| Scheduler | Go runtime (user space) | OS kernel |
| Switching | cheap | requires syscall |

You can easily run **hundreds of thousands** of goroutines. This is why Go excels at concurrent I/O.

## The GMP scheduler

Go uses an **M:N** scheduler:

- **G** — goroutine (your code + stack).
- **M** — machine (OS thread).
- **P** — processor (a context for scheduling, usually `GOMAXPROCS()` of them).

The runtime multiplexes many Gs onto a few Ms. When a goroutine blocks on I/O, the scheduler parks it and runs another, so threads stay busy.

Set parallelism:

```go
runtime.GOMAXPROCS(4)   // default: number of logical CPUs
```

## When goroutines are useful

- Handling many network connections concurrently.
- Parallelizing independent CPU work (subject to `GOMAXPROCS`).
- Pipelines of processing stages.
- Background tasks (with cancellation — Chapter 18).

## Waiting with `sync.WaitGroup`

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id)
    }(i)
}

wg.Wait()
fmt.Println("all done")
```

Rules:
- Call `Add` **before** starting the goroutine (not inside it).
- Always `defer wg.Done()` so it runs even on panic.
- `Wait` blocks until the counter reaches zero.

## Common pitfalls

### 1. `main` exits before goroutines finish

```go
go doWork()
// main returns → program exits → doWork never completes
```

Wait using a `WaitGroup`, channel, or `errgroup`.

### 2. Loop variable capture (pre-Go 1.22)

```go
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()   // old: often prints 3 3 3
}
```

Go 1.22+ makes loop variables per-iteration. For older Go, pass as a parameter: `go func(i int){...}(i)`.

### 3. Data races

Concurrent reads/writes to shared memory without synchronization are a **data race**:

```go
counter := 0
for i := 0; i < 1000; i++ {
    go func() { counter++ }()   // RACE!
}
```

Fix with `sync.Mutex`, atomics (`sync/atomic`), or channels. Always run tests with the race detector:

```bash
go test -race ./...
go run -race .
```

### 4. Goroutine leaks

A goroutine that blocks forever on a channel, mutex, or network call is never collected. Always provide a way to stop (cancellation via `context`, Chapter 18).

### 5. Starting goroutines in a loop with unbounded work

Limit concurrency with a worker pool or semaphore (Chapter 18) — otherwise you may exhaust memory or file descriptors.

## Goroutines and closures

Capturing variables is fine if you own them per goroutine:

```go
for _, url := range urls {
    url := url   // explicit copy (needed pre-1.22)
    go func() {
        fetch(url)
    }()
}
```

## Printing interleaves

Goroutines write to stdout concurrently; output may interleave. Use `fmt.Println` (each call is fairly atomic) or a mutex if you need grouped output.

## Runnable example

See `examples/15-goroutines/main.go`.

## Chapter summary

- `go f()` starts a goroutine; cheap, dynamically sized stack.
- The Go runtime schedules goroutines onto OS threads (GMP model).
- Use `sync.WaitGroup` to wait: `Add` before, `defer Done` inside, `Wait` after.
- Go 1.22 gives loop variables per-iteration semantics.
- Always check for data races with `-race`; prevent leaks with cancellation.

## Exercises

1. Launch 10 goroutines that print their ID and wait with a `WaitGroup`.
2. Create a data race intentionally, then fix it with a mutex.
3. Measure how many goroutines you can start before memory pressure becomes an issue.
4. Show that output from two goroutines can interleave.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 16)*
