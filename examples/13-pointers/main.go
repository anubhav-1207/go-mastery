// Demonstrates addresses, dereferencing, and pointer semantics.
package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

type Person struct{ Name string }

func rename(p *Person, name string) { p.Name = name }

func byValue(n int)    { n = 99 }
func byPointer(n *int) { *n = 99 }

func main() {
	x := 42
	p := &x
	fmt.Println("x:", x, "address:", p, "value at p:", *p)
	*p = 100
	fmt.Println("x after *p=100:", x)

	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr, "is nil:", nilPtr == nil)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("swapped:", a, b)

	person := &Person{Name: "Alice"}
	rename(person, "Bob")
	fmt.Println("renamed:", person.Name)

	v := 1
	byValue(v)
	fmt.Println("after byValue:", v)
	byPointer(&v)
	fmt.Println("after byPointer:", v)

	// Slices/maps are descriptors: no pointer needed to mutate elements
	s := []int{1, 2, 3}
	m := map[string]int{"a": 1}
	mutate(s, m)
	fmt.Println("slice/map:", s, m)

	// new() allocates a zeroed value
	n := new(int)
	fmt.Println("new(int):", *n)
}

func mutate(s []int, m map[string]int) {
	s[0] = 99
	m["b"] = 2
}
