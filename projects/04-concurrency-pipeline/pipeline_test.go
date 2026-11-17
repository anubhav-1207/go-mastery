package main

import (
	"context"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	jobs := []Job{1, 2, 3, 4, 5}
	results := Run(context.Background(), jobs, 3)
	if len(results) != len(jobs) {
		t.Fatalf("got %d results; want %d", len(results), len(jobs))
	}

	seen := map[Job]int{}
	for _, r := range results {
		if r.Squared != int(r.Job)*int(r.Job) {
			t.Errorf("job %d: squared = %d", r.Job, r.Squared)
		}
		seen[r.Job]++
	}
	for _, j := range jobs {
		if seen[j] != 1 {
			t.Errorf("job %d processed %d times", j, seen[j])
		}
	}
}

func TestRunCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	jobs := make([]Job, 1000)
	results := Run(ctx, jobs, 2)
	if len(results) == len(jobs) {
		t.Error("expected cancellation to stop early")
	}
}
