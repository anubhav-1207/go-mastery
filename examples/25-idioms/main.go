// Demonstrates common Go idioms: comma-ok, zero values, compile-time checks,
// guard clauses, functional options, and blank identifiers.
package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type Notifier interface{ Notify(string) error }

type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(msg string) error {
	fmt.Println("notify:", msg)
	return nil
}

// Compile-time interface check
var _ Notifier = ConsoleNotifier{}

type Server struct {
	addr    string
	timeout time.Duration
}

type Option func(*Server)

func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }

func NewServer(addr string, opts ...Option) *Server {
	s := &Server{addr: addr, timeout: 5 * time.Second}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func process(m map[string]int, key string) error {
	// Guard clause
	if key == "" {
		return fmt.Errorf("empty key")
	}
	// comma-ok
	if value, ok := m[key]; ok {
		fmt.Println("found:", value)
	} else {
		return fmt.Errorf("missing %q", key)
	}
	return nil
}

func main() {
	// Useful zero values
	var buf bytes.Buffer
	buf.WriteString("zero value is usable")
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Wait()
	var s []int
	fmt.Println(buf.String(), "wg zero ok, slice nil:", s == nil)

	m := map[string]int{"a": 1}
	_ = process(m, "a")
	fmt.Println("process missing:", process(m, "x"))

	// Blank identifier ignores index
	for _, v := range []int{1, 2, 3} {
		_ = v
	}

	srv := NewServer(":8080", WithTimeout(time.Second))
	fmt.Printf("server: %+v\n", srv)

	// Expressionless switch as an if/else chain
	n := 5
	switch {
	case n < 0:
		fmt.Println("negative")
	case n == 0:
		fmt.Println("zero")
	default:
		fmt.Println("positive")
	}
}
