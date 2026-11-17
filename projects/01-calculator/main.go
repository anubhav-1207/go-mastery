package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	a := flag.Float64("a", 0, "first number")
	b := flag.Float64("b", 0, "second number")
	op := flag.String("op", "add", "operation: add, sub, mul, div, pow")
	flag.Parse()

	result, err := Calculate(*op, *a, *b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("%g %s %g = %g\n", *a, *op, *b, result)
}
