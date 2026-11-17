// Demonstrates functions, multiple returns, variadics, closures, and recursion.
package main

import "fmt"

func add(a, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return (short function — acceptable)
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func mapInts(nums []int, fn func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = fn(n)
	}
	return out
}

func main() {
	fmt.Println("add:", add(2, 3))

	result, err := divide(10, 2)
	fmt.Println("divide:", result, err)
	_, err = divide(10, 0)
	fmt.Println("divide by zero err:", err)

	x, y := split(10)
	fmt.Println("split:", x, y)

	fmt.Println("sum:", sum(), sum(1, 2, 3))
	values := []int{4, 5, 6}
	fmt.Println("sum spread:", sum(values...))

	c := counter()
	fmt.Println(c(), c(), c())

	fmt.Println("factorial(5):", factorial(5))

	doubled := mapInts([]int{1, 2, 3}, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	// Function value
	f := add
	fmt.Println("function value:", f(7, 8))
}
