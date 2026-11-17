// Package main is our very first Go program.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, world!")

	name := "Gopher"
	age := 7
	fmt.Printf("Name: %s, age: %d\n", name, age)
	fmt.Printf("Default format: %v, type: %T\n", name, name)
	fmt.Printf("Quoted: %q\n", name)

	fmt.Println("Running on", runtime.GOOS, runtime.GOARCH)
}
