// A tiny multi-package module demonstrating exports and imports.
package main

import (
	"fmt"

	"github.com/anubhav-1207/go-mastery/examples/14-packages/mathutil"
)

func main() {
	fmt.Println("Add:", mathutil.Add(2, 3))
	fmt.Println("Area:", mathutil.Rectangle{Width: 3, Height: 4}.Area())
	fmt.Println("Pi:", mathutil.Pi)
	fmt.Println("Version:", mathutil.Version())
}
