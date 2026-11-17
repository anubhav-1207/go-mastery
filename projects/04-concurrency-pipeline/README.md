# Project 04 — Concurrency Pipeline

**Skills:** goroutines, channels, fan-out/fan-in, `sync.WaitGroup`, `context` cancellation, `select`.

## Run

```bash
go run ./projects/04-concurrency-pipeline -jobs 12 -workers 4
go run ./projects/04-concurrency-pipeline -jobs 12 -workers 4 -timeout 1ms
```

## Test

```bash
go test -race ./projects/04-concurrency-pipeline
```

## What to notice

- `generate` owns the job channel and closes it; `worker`s never close the shared result channel.
- A dedicated goroutine calls `wg.Wait()` then closes `resultCh`, so `range` terminates cleanly.
- Every blocking operation is inside a `select` with `ctx.Done()`, so no goroutine leaks.
- The race detector should report nothing.
