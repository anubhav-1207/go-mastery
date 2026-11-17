// Demonstrates implicit interfaces, type assertions, type switches, and io.Reader.
package main

import (
	"fmt"
	"io"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (Cat) Speak() string { return "Meow!" }

type Notifier interface{ Notify(string) error }

type EmailNotifier struct{}

func (EmailNotifier) Notify(msg string) error {
	fmt.Println("email:", msg)
	return nil
}

func makeItSpeak(s Speaker) { fmt.Println(s.Speak()) }

func describe(v any) {
	switch x := v.(type) {
	case int:
		fmt.Println("int:", x*2)
	case string:
		fmt.Println("string:", strings.ToUpper(x))
	case nil:
		fmt.Println("nil")
	default:
		fmt.Printf("unknown type %T\n", x)
	}
}

func main() {
	makeItSpeak(Dog{})
	makeItSpeak(Cat{})

	// Compile-time interface check
	var _ Speaker = Dog{}
	var _ Notifier = EmailNotifier{}

	// Type assertion
	var x any = "hello"
	s, ok := x.(string)
	fmt.Println("assert:", s, ok)
	_, ok = x.(int)
	fmt.Println("int assertion:", ok)

	describe(10)
	describe("hi")
	describe(3.14)
	describe(nil)

	// io.Reader / io.Writer
	var r io.Reader = strings.NewReader("hello reader")
	data := make([]byte, 5)
	n, err := r.Read(data)
	fmt.Println("read:", string(data[:n]), err)

	// Typed nil pitfall
	var d *Dog
	var speaker Speaker = d
	fmt.Println("typed nil interface is nil?", speaker == nil)
}
