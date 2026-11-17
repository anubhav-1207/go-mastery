// Demonstrates variable declarations, zero values, constants, and iota.
package main

import "fmt"

const (
	Sunday  = iota // 0
	Monday         // 1
	Tuesday        // 2
)

const (
	_  = iota             // discard 0
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)

func main() {
	// Zero values
	var count int
	var text string
	var active bool
	var numbers []int
	fmt.Printf("zero values: %d %q %t %v (nil=%t)\n",
		count, text, active, numbers, numbers == nil)

	// Explicit and inferred declarations
	var age int = 30
	name := "Alice"
	fmt.Println(name, age)

	// Grouped declarations
	var (
		x int     = 1
		y float64 = 2.5
	)
	fmt.Println(x, y)

	// Short redeclaration: at least one new variable
	a := 10
	b := 20
	a, c := 99, "new"
	fmt.Println(a, b, c)

	// Untyped constants adapt to the required type
	const n = 10
	var i int = n
	var f float64 = n
	fmt.Println(i, f)

	fmt.Println(Sunday, Monday, Tuesday)
	fmt.Println(KB, MB, GB)

	// Shadowing
	outer := 1
	{
		outer := 2
		fmt.Println("inner:", outer)
	}
	fmt.Println("outer:", outer)
}
