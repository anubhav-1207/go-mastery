// Demonstrates channels: send/receive, close, range, direction, fan-out, fan-in.
package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(out chan<- int, n int) {
	defer close(out)
	for i := 0; i < n; i++ {
		out <- i
	}
}

func square(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		out <- v * v
	}
}

func main() {
	// Basic unbuffered channel
	ch := make(chan string)
	go func() { ch <- "ping" }()
	fmt.Println(<-ch)

	// Buffered channel
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println(<-buffered, <-buffered)

	// Closed channel + comma-ok
	c := make(chan int, 2)
	c <- 10
	c <- 20
	close(c)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println("value:", v)
	}

	// Range over channel
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	nums <- 3
	close(nums)
	for v := range nums {
		fmt.Println("ranged:", v)
	}

	// Pipeline with fan-out (2 workers) and fan-in
	jobs := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup
	for w := 0; w < 2; w++ {
		wg.Add(1)
		go square(jobs, squares, &wg)
	}
	go func() { wg.Wait(); close(squares) }()

	producer(jobs, 6)

	for v := range squares {
		fmt.Println("square:", v)
	}

	// Timeout with select
	slow := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		slow <- "done"
	}()
	select {
	case msg := <-slow:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
