// Demonstrates select, context cancellation, worker pools, errgroup-style errors.
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}

func main() {
	// select with timeout
	ch := make(chan string)
	go func() { time.Sleep(30 * time.Millisecond); ch <- "result" }()
	select {
	case v := <-ch:
		fmt.Println("got:", v)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("timed out")
	}

	// context cancellation
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine cancelled")
				return
			default:
				time.Sleep(5 * time.Millisecond)
			}
		}
	}()
	time.Sleep(20 * time.Millisecond)
	cancel()
	<-done

	// context deadline
	ctx2, cancel2 := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel2()
	<-ctx2.Done()
	fmt.Println("deadline:", ctx2.Err())

	// Worker pool
	jobs := make(chan int, 20)
	results := make(chan int, 20)
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
		fmt.Print(r, " ")
	}
	fmt.Println()

	// Mutex + atomic
	var mu sync.Mutex
	var lockedCounter int
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			lockedCounter++
			mu.Unlock()
		}()
	}
	var atomicCounter atomic.Int64
	for i := 0; i < 100; i++ {
		go func() { atomicCounter.Add(1) }()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("mutex:", lockedCounter, "atomic:", atomicCounter.Load())

	// once
	var once sync.Once
	once.Do(func() { fmt.Println("initialized once") })
	once.Do(func() { fmt.Println("this never prints") })

	// error aggregation
	errs := []error{nil, errors.New("a failed"), nil, errors.New("b failed")}
	var first error
	for _, e := range errs {
		if e != nil && first == nil {
			first = e
		}
	}
	fmt.Println("first error:", first)
}
