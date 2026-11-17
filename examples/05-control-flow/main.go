// Demonstrates if, for, range, switch, labels, and defer.
package main

import "fmt"

func classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	case n < 10:
		return "small positive"
	default:
		return "large positive"
	}
}

func main() {
	// if with init statement
	if err := fmt.Errorf("sample"); err != nil {
		fmt.Println("handled:", err)
	}

	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// Loop forms
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	n := 3
	for n > 0 {
		n--
	}
	fmt.Println("n =", n)

	// range over slice, string, map
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	for index, r := range "Go" {
		fmt.Printf("(%d:%c) ", index, r)
	}
	fmt.Println()

	m := map[string]int{"a": 1}
	for k, v := range m {
		fmt.Printf("%s=%d ", k, v)
	}
	fmt.Println()

	// Labeled break
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j == 4 {
				fmt.Println("breaking at", i, j)
				break Outer
			}
		}
	}

	fmt.Println(classify(-5), classify(0), classify(5), classify(50))

	// fallthrough
	switch 1 {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
	}

	// defer runs LIFO
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")
	fmt.Println("in main")
}
