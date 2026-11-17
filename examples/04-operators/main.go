// Demonstrates arithmetic, comparison, logical, and bitwise operators.
package main

import (
	"fmt"
	"strings"
)

const (
	Read    = 1 << iota // 1
	Write               // 2
	Execute             // 4
)

func main() {
	fmt.Println("7/2 =", 7/2, "7%2 =", 7%2)
	fmt.Println("7/2.0 =", 7/2.0)
	fmt.Println("-7/2 =", -7/2, "-7%2 =", -7%2)

	i := 0
	i++
	i += 5
	fmt.Println("i =", i)

	fmt.Println("'Z' < 'a':", "Z" < "a")

	// Short-circuit: 10/x is not evaluated when x == 0
	x := 0
	if x != 0 && 10/x > 2 {
		fmt.Println("big")
	} else {
		fmt.Println("short-circuit kept us safe")
	}

	// Bitwise operators
	a := 0b1100 // 12
	b := 0b1010 // 10
	fmt.Printf("a=%04b b=%04b\n", a, b)
	fmt.Printf("a&b =%04b (%d)\n", a&b, a&b)
	fmt.Printf("a|b =%04b (%d)\n", a|b, a|b)
	fmt.Printf("a^b =%04b (%d)\n", a^b, a^b)
	fmt.Printf("a&^b=%04b (%d)\n", a&^b, a&^b)
	fmt.Println("1<<10 =", 1<<10, "1024>>3 =", 1024>>3)
	fmt.Println("^uint8(0) =", ^uint8(0))

	// Flags
	perm := Read | Write
	fmt.Printf("perm=%d read=%t execute=%t\n",
		perm, perm&Read != 0, perm&Execute != 0)

	// String building
	var sb strings.Builder
	for n := 0; n < 5; n++ {
		sb.WriteString("item ")
		sb.WriteString(fmt.Sprint(n))
		sb.WriteByte('\n')
	}
	fmt.Print(sb.String())
}
