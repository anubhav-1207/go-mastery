package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

func main() {
	count := flag.Int("jobs", 20, "number of jobs")
	workers := flag.Int("workers", 4, "number of workers")
	timeout := flag.Duration("timeout", 0, "optional timeout (e.g. 1ms)")
	flag.Parse()

	jobs := make([]Job, *count)
	for i := range jobs {
		jobs[i] = Job(i + 1)
	}

	ctx := context.Background()
	if *timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	start := time.Now()
	results := Run(ctx, jobs, *workers)
	fmt.Print(FormatResults(results))
	fmt.Printf("processed %d/%d jobs in %s\n", len(results), len(jobs), time.Since(start))
}
