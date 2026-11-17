// Demonstrates value/pointer receivers, method promotion, and Stringer.
package main

import (
	"fmt"
	"strconv"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r *Rectangle) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

type Counter struct{ n int }

func (c *Counter) Inc()       { c.n++ }
func (c *Counter) Value() int { return c.n }

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}
	return strconv.Itoa(n.Value)
}

type Reader struct{}

func (Reader) Read() string { return "data" }

type File struct {
	Reader // embedded
	Name   string
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println("area:", r.Area(), "perimeter:", r.Perimeter())
	r.Scale(2)
	fmt.Println("scaled area:", r.Area())

	c := &Counter{}
	c.Inc()
	c.Inc()
	fmt.Println("counter:", c.Value())

	var n *Node
	fmt.Println("nil node:", n.String())

	f := File{Name: "notes.txt"}
	fmt.Println("promoted:", f.Read())

	// Methods on named non-struct types
	type MyInt int
	fmt.Println(MyInt(5) * 2)

	// fmt.Stringer integration
	fmt.Println("node:", &Node{Value: 42})
}
