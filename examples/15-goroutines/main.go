// Demonstrates goroutines, WaitGroup, and the importance of synchronization.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printNumbers(prefix string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(prefix, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	// Fire and wait with WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("worker %d starting\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("worker %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("all workers finished")

	// Loop variable capture (Go 1.22+ gives per-iteration variables)
	for _, word := range []string{"a", "b", "c"} {
		go func(w string) {
			fmt.Println("word:", w)
		}(word)
	}
	time.Sleep(50 * time.Millisecond)

	// Demonstrate a data race conceptually (commented safe version provided)
	// Unsafe: shared++ from many goroutines without synchronization.
	// Safe version with a mutex:
	var mu sync.Mutex
	shared := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			shared++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("shared counter:", shared)

	// Interleaving output
	go printNumbers("A", 3)
	printNumbers("B", 3)
}
