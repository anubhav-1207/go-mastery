package main

import (
	"context"
	"fmt"
	"sync"
)

// Job is a unit of work: a number to process.
type Job int

// Result holds the processed value and which worker produced it.
type Result struct {
	Job     Job
	Squared int
	Worker  int
}

// generate sends jobs on out until the input is exhausted or ctx is cancelled.
// The sender owns out and always closes it.
func generate(ctx context.Context, jobs []Job, out chan<- Job) {
	defer close(out)
	for _, j := range jobs {
		select {
		case <-ctx.Done():
			return
		case out <- j:
		}
	}
}

// worker processes jobs until the input closes or ctx is cancelled.
// It never closes out (multiple workers share it) — the fan-in goroutine does.
func worker(ctx context.Context, id int, in <-chan Job, out chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-in:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- Result{Job: j, Squared: int(j) * int(j), Worker: id}:
			}
		}
	}
}

// Run executes a fan-out/fan-in worker-pool pipeline.
// It returns the results collected before completion or cancellation.
func Run(ctx context.Context, jobs []Job, numWorkers int) []Result {
	if numWorkers < 1 {
		numWorkers = 1
	}

	jobCh := make(chan Job)
	resultCh := make(chan Result)

	go generate(ctx, jobs, jobCh)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobCh, resultCh, &wg)
	}

	// Fan-in closer: close resultCh only after every worker is done.
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make([]Result, 0, len(jobs))
	for r := range resultCh {
		results = append(results, r)
	}
	return results
}

// FormatResults renders results as a human-readable report.
func FormatResults(results []Result) string {
	out := ""
	for _, r := range results {
		out += fmt.Sprintf("job %d squared = %d (worker %d)\n", r.Job, r.Squared, r.Worker)
	}
	return out
}
