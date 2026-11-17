# Chapter 24 — Performance & Optimization

> "Premature optimization is the root of all evil." — Donald Knuth

Write clear code first, **measure**, then optimize the bottleneck.

## Benchmarking

```go
func BenchmarkConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = "a" + "b" + "c"
    }
}
```

```bash
go test -bench=. -benchmem -count=3 ./...
```

Look for:
- `ns/op` — latency per operation.
- `B/op` — bytes allocated.
- `allocs/op` — allocations per operation.

Fewer allocations usually mean less GC pressure and better latency.

## Profiling

### CPU profile

```bash
go test -cpuprofile=cpu.out ./...
go tool pprof cpu.out
# inside pprof: top, list functionName, web
```

Or in code:

```go
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

### Memory profile

```bash
go test -memprofile=mem.out ./...
go tool pprof mem.out
```

### Execution tracer

```bash
go test -trace=trace.out ./...
go tool trace trace.out
```

The tracer shows goroutine scheduling, blocking, and GC events over time — invaluable for concurrency issues.

### Escape analysis

```bash
go build -gcflags='-m -m' .
```

Look for `escapes to heap`. Reducing escapes cuts allocations.

## Common optimizations

### 1. Preallocate slices

```go
// Bad: grows repeatedly
var out []int
for _, v := range in {
    out = append(out, v)
}

// Good: one allocation
out := make([]int, 0, len(in))
```

### 2. Use `strings.Builder` for concatenation

```go
var sb strings.Builder
sb.Grow(estimatedSize)
for _, part := range parts {
    sb.WriteString(part)
}
```

### 3. Avoid unnecessary allocations

- Pass large structs by pointer (or use pointers in slices).
- Reuse buffers with `sync.Pool`:

```go
var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

buf := bufPool.Get().(*bytes.Buffer)
buf.Reset()
defer bufPool.Put(buf)
```

### 4. Prefer arrays / fixed sizes in hot loops when possible.

### 5. Avoid `interface{}`/reflection in hot paths

`json.Marshal` with reflection is slower than generated code. Measure.

### 6. Batch I/O

Use `bufio.Writer`/`Reader` and avoid one syscall per item.

### 7. Minimize locking

Shard data, use `RWMutex`, or use atomics for simple counters.

### 8. Be careful with goroutine overhead

For tiny CPU tasks, spawning a goroutine costs more than doing it inline. Use worker pools for many small tasks.

## Garbage collection

Go uses a concurrent, tri-color mark-and-sweep collector. Tuning is rarely needed, but:

```bash
GOGC=200 go run .     # collect less often (higher throughput, more memory)
GOMEMLIMIT=1GiB go run .
```

Reducing allocation rate is almost always better than tuning GC.

## Compiler optimizations

Go's compiler inlines small functions, eliminates bounds checks when it can prove safety, and does escape analysis. Write straightforward code and let the compiler help. You can inspect inlining:

```bash
go build -gcflags='-m' .
```

## `testing.AllocsPerRun`

```go
allocs := testing.AllocsPerRun(100, func() { doWork() })
```

Useful in tests to guard against accidental allocation regressions.

## When to stop optimizing

Optimize when:
- Profiling shows a real hotspot.
- Latency/throughput requirements are not met.
- Allocations cause GC pauses affecting SLOs.

Do not optimize based on intuition alone.

## Runnable example

See `examples/23-performance/` (benchmarks + profiling notes).

## Chapter summary

- Measure with benchmarks (`-benchmem`) and profiles (`pprof`, `trace`).
- Reduce allocations: preallocate, use builders, reuse buffers.
- Escape analysis explains heap allocations.
- Batch I/O and bound concurrency.
- Optimize only what profiling proves is hot.

## Exercises

1. Benchmark `append` with and without preallocation.
2. Profile a slow function with `pprof` and identify the hotspot.
3. Use `testing.AllocsPerRun` to verify a function is allocation-free.
4. Compare `json.Marshal` with a hand-written encoder for a hot struct.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 24)*
