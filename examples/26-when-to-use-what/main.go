// Side-by-side demonstrations of common Go choices.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Slice: ordered sequence
	sequence := []int{3, 1, 2}
	fmt.Println("slice:", sequence)

	// Map: key lookup
	lookup := map[string]int{"one": 1, "two": 2}
	fmt.Println("map lookup:", lookup["two"])

	// Set: map[T]struct{}
	set := map[string]struct{}{"x": {}}
	_, exists := set["x"]
	fmt.Println("set:", exists)

	// Pointer receiver for mutation
	type Counter struct{ n int }
	inc := func(c *Counter) { c.n++ }
	c := &Counter{}
	inc(c)
	fmt.Println("counter:", c.n)

	// Channel for passing data
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println("channel:", <-ch)

	// Mutex for protecting shared state
	var mu sync.Mutex
	shared := 0
	mu.Lock()
	shared++
	mu.Unlock()
	fmt.Println("shared:", shared)

	// error for expected failure
	if err := mightFail(false); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("error: nil (expected failure handled)")
	}
}

func mightFail(fail bool) error {
	if fail {
		return fmt.Errorf("something went wrong")
	}
	return nil
}
