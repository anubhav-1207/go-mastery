// Demonstrates maps, the comma-ok idiom, sets, and deterministic iteration.
package main

import (
	"fmt"
	"slices"
)

func main() {
	ages := map[string]int{"Alice": 30, "Bob": 25}
	ages["Charlie"] = 35
	fmt.Println("ages:", ages)

	fmt.Println("Alice:", ages["Alice"])
	fmt.Println("Unknown:", ages["Unknown"])

	// comma-ok
	if age, ok := ages["Bob"]; ok {
		fmt.Println("Bob's age:", age)
	}
	if _, ok := ages["Nobody"]; !ok {
		fmt.Println("Nobody not found")
	}

	delete(ages, "Bob")
	fmt.Println("after delete:", ages)

	// Set using struct{} values
	set := map[string]struct{}{"a": {}, "b": {}}
	set["c"] = struct{}{}
	_, present := set["b"]
	fmt.Println("set has b:", present, "len:", len(set))
	delete(set, "a")

	// Preallocation
	m := make(map[int]string, 100)
	m[1] = "one"

	// Deterministic iteration: collect and sort keys
	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Printf("  %s=%d\n", k, ages[k])
	}

	// Clear
	clear(ages)
	fmt.Println("after clear:", len(ages))

	// Grouping
	words := []string{"apple", "avocado", "banana", "blueberry", "apricot"}
	groups := map[string][]string{}
	for _, w := range words {
		first := string(w[0])
		groups[first] = append(groups[first], w)
	}
	fmt.Println("groups:", groups)
}
