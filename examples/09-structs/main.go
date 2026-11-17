// Demonstrates structs, tags, embedding, and value semantics.
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string
	Zip  string
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
	Secret  string  `json:"-"`
}

type Animal struct {
	Name string
}

func (a Animal) Speak() string { return "..." }

type Dog struct {
	Animal // embedded
	Breed  string
}

func main() {
	p := Person{
		Name:    "Alice",
		Age:     30,
		Address: Address{City: "Paris", Zip: "75001"},
		Secret:  "hidden",
	}
	fmt.Println(p.Name, p.Address.City)

	// Value semantics
	q := p
	q.Age = 40
	fmt.Println("p.Age:", p.Age, "q.Age:", q.Age)

	// Pointer semantics
	r := &p
	r.Age = 50
	fmt.Println("p.Age after pointer:", p.Age)

	// JSON with tags
	b, _ := json.Marshal(p)
	fmt.Println("json:", string(b))

	// Anonymous struct
	point := struct{ X, Y int }{X: 1, Y: 2}
	fmt.Println("point:", point)

	// Embedding and promotion
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	fmt.Println("dog:", d.Name, d.Breed, d.Speak())

	// Comparable struct
	type Pt struct{ X, Y int }
	fmt.Println("Pt equal:", Pt{1, 2} == Pt{1, 2})

	var zero Person
	fmt.Printf("zero: %+v\n", zero)
}
