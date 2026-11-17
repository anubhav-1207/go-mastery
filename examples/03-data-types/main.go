// Demonstrates Go's basic data types, runes, and conversions.
package main

import (
	"fmt"
	"strconv"
)

type Age int          // a distinct named type
type ByteAlias = byte // alias for uint8

func main() {
	// Integers
	var i int = 42
	var u uint8 = 200
	var b byte = 'A'
	fmt.Println(i, u, b, string(b))

	// Floats
	var pi float64 = 3.14159
	fmt.Printf("%.3f\n", pi)

	// Booleans
	var ok bool = true
	fmt.Println(!ok)

	// Complex
	c := complex(3, 4)
	fmt.Println(c, real(c), imag(c))

	// Strings: bytes vs runes
	s := "Hello, 世界"
	fmt.Println("bytes:", len(s), "runes:", len([]rune(s)))
	for index, r := range s {
		fmt.Printf("  [%d] %c U+%04X\n", index, r, r)
	}

	// Concatenation and conversion
	fmt.Println("a" + "b" + strconv.Itoa(123))
	fmt.Println(string([]byte{'g', 'o'}))

	// Explicit conversions (truncation/wrapping).
	// Use variables: converting an untyped constant that would overflow is a compile error.
	f := 3.99
	big := 300
	fmt.Println("int(3.99):", int(f))
	fmt.Println("uint8(300):", uint8(big))

	// Named type needs conversion
	var age Age = Age(30)
	fmt.Println(age)

	// Alias needs none
	var alias ByteAlias = 255
	fmt.Println(alias)

	// Reverse a string rune-by-rune
	runes := []rune("hello 世界")
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	fmt.Println("reversed:", string(runes))
}
