// Demonstrates arrays, slices, append, copy, and the slices package.
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Arrays are values
	var a [3]int
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3, 4}
	d := [5]int{2: 42}
	fmt.Println("arrays:", a, b, c, d)

	x := [2]int{1, 2}
	y := x
	y[0] = 99
	fmt.Println("array copy:", x, y)

	// Slices: len and cap
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4]
	fmt.Printf("slice %v len=%d cap=%d\n", s, len(s), cap(s))

	sliced := make([]int, 3, 5)
	fmt.Printf("make: len=%d cap=%d\n", len(sliced), cap(sliced))

	// append (always assign the result)
	sliced = append(sliced, 4)
	sliced = append(sliced, 5, 6)
	fmt.Println("appended:", sliced)

	other := []int{7, 8}
	sliced = append(sliced, other...)
	fmt.Println("spread append:", sliced)

	// Slices share backing arrays
	p := []int{1, 2, 3}
	q := p
	q[0] = 99
	fmt.Println("shared:", p, q)

	// copy to detach
	detached := make([]int, len(p))
	copy(detached, p)
	detached[0] = 0
	fmt.Println("detached:", p, detached)

	// nil vs empty
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nil=%t empty=%t lens=%d,%d\n",
		nilSlice == nil, emptySlice == nil, len(nilSlice), len(emptySlice))

	// Remove element at index i
	nums := []int{1, 2, 3, 4, 5}
	i := 2
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("after remove:", nums)

	// slices package
	values := []int{3, 1, 2, 1}
	slices.Sort(values)
	fmt.Println("sorted:", values)
	fmt.Println("contains 2:", slices.Contains(values, 2))
	fmt.Println("index of 3:", slices.Index(values, 3))
	fmt.Println("compact:", slices.Compact(values))
	cloned := slices.Clone(values)
	cloned[0] = 100
	fmt.Println("clone:", values, cloned)

	// 2D slice
	grid := make([][]int, 2)
	for r := range grid {
		grid[r] = make([]int, 3)
		for c := range grid[r] {
			grid[r][c] = r*3 + c
		}
	}
	fmt.Println("grid:", grid)
}
