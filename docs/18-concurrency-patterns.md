# Chapter 18 — Concurrency Patterns

## `select`

`select` waits on multiple communication operations simultaneously. It blocks until one can proceed.

```go
select {
case msg := <-messages:
    fmt.Println("received", msg)
case sig := <-signals:
    fmt.Println("signal", sig)
default:
    fmt.Println("no activity")   // non-blocking
}
```

Rules:
- Each `case` must be a channel send/receive.
- If multiple are ready, one is chosen **randomly**.
- A `default` makes it non-blocking.
- An empty `select {}` blocks forever (used to keep `main` alive).

### Timeouts

```go
select {
case result := <-doWork():
    fmt.Println(result)
case <-time.After(2 * time.Second):
    fmt.Println("timed out")
}
```

## `context` — cancellation and deadlines

`context.Context` carries deadlines, cancellation signals, and request-scoped values across APIs.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // do a unit of work
        }
    }
}()

time.Sleep(time.Second)
cancel()   // goroutine observes ctx.Done() and exits
```

### Variants

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
ctx := context.WithValue(parent, key, value)   // request-scoped data (use sparingly)
```

Always `defer cancel()` to release resources, even for timeouts.

### Convention

Pass `ctx` as the **first parameter** of functions that do I/O:

```go
func Fetch(ctx context.Context, url string) ([]byte, error)
```

Check `ctx.Err()` in loops and propagate it.

## `sync` package essentials

### Mutex

```go
var mu sync.Mutex
var counter int

func inc() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

Always `defer Unlock` so panics do not leave it locked. `sync.Mutex` must not be copied after first use — use a pointer in structs.

### RWMutex

Allows multiple concurrent readers or one writer:

```go
var mu sync.RWMutex

mu.RLock()
value := data[key]
mu.RUnlock()

mu.Lock()
data[key] = value
mu.Unlock()
```

### WaitGroup

Wait for a collection of goroutines (Chapter 16).

### Once

Run initialization exactly once:

```go
var once sync.Once

func setup() { once.Do(initialize) }
```

### Atomic operations

For simple counters, `sync/atomic` is lighter than a mutex:

```go
var counter atomic.Int64
counter.Add(1)
fmt.Println(counter.Load())
```

### `sync.Map`

Specialized concurrent map optimized for read-mostly / disjoint-key workloads. For most cases, a plain map + `sync.RWMutex` is clearer.

## Worker pool

Limit concurrency and reuse goroutines:

```go
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        results <- j * 2
    }
}

jobs := make(chan int, 100)
results := make(chan int, 100)

var wg sync.WaitGroup
for w := 1; w <= 3; w++ {
    wg.Add(1)
    go worker(w, jobs, results, &wg)
}

for j := 1; j <= 9; j++ {
    jobs <- j
}
close(jobs)

go func() { wg.Wait(); close(results) }()
for r := range results {
    fmt.Println(r)
}
```

## Rate limiting with a ticker

```go
ticker := time.NewTicker(100 * time.Millisecond)
defer ticker.Stop()

for job := range jobs {
    <-ticker.C        // at most 10 jobs/second
    process(job)
}
```

## Semaphore pattern

Cap concurrent operations with a buffered channel:

```go
sem := make(chan struct{}, 10)   // max 10 at a time

for _, task := range tasks {
    sem <- struct{}{}            // acquire
    go func(t Task) {
        defer func() { <-sem }() // release
        run(t)
    }(task)
}
```

## `errgroup` (golang.org/x/sync/errgroup)

Run goroutines and collect the first error:

```go
g, ctx := errgroup.WithContext(context.Background())

for _, url := range urls {
    url := url
    g.Go(func() error {
        _, err := fetch(ctx, url)
        return err
    })
}

if err := g.Wait(); err != nil {
    return err
}
```

`errgroup.WithContext` cancels the shared context when any goroutine returns an error.

## Pipeline with cancellation

Combine channels, context, and `select`:

```go
func stage(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-ctx.Done():
                return
            case v, ok := <-in:
                if !ok { return }
                out <- v * 2
            }
        }
    }()
    return out
}
```

## Race detector

```bash
go test -race ./...
```

The race detector instruments memory accesses and reports unsynchronized concurrent access. Use it in CI.

## Runnable example

See `examples/17-concurrency-patterns/main.go`.

## Chapter summary

- `select` multiplexes channel operations; `default` makes it non-blocking.
- `context` propagates cancellation/deadlines — pass it first, always call `cancel()`.
- Use `Mutex`/`RWMutex` for shared state, `atomic` for simple counters.
- Worker pools and semaphores bound concurrency.
- `errgroup` collects errors and cancels siblings.
- Always test with `-race`.

## Exercises

1. Write a `select` with a 500 ms timeout.
2. Implement a worker pool with 4 workers and 20 jobs.
3. Use `context.WithTimeout` to cancel a long-running goroutine.
4. Build an `errgroup` that fetches multiple URLs and fails on the first error.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 18)*
