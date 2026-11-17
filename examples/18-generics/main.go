// Demonstrates generic functions, constraints, and generic types.
package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	~int | ~int64 | ~float64
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Sum[T Number](xs []T) (total T) {
	for _, x := range xs {
		total += x
	}
	return
}

func Map[T any, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](in []T, pred func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	n := len(s.items) - 1
	v := s.items[n]
	s.items = s.items[:n]
	return v, true
}

type Age int

func main() {
	fmt.Println("Max:", Max(3, 7), Max("a", "b"), Max(1.5, 2.5))
	fmt.Println("Sum ints:", Sum([]int{1, 2, 3}))
	fmt.Println("Sum floats:", Sum([]float64{1.5, 2.5}))
	fmt.Println("Sum Ages:", Sum([]Age{1, 2, 3}))

	strs := Map([]int{1, 2, 3}, func(n int) string {
		return fmt.Sprintf("n%d", n)
	})
	fmt.Println("Map:", strs)

	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filter:", evens)

	stack := &Stack[string]{}
	stack.Push("a")
	stack.Push("b")
	v, ok := stack.Pop()
	fmt.Println("Pop:", v, ok)
	v, ok = stack.Pop()
	fmt.Println("Pop:", v, ok)
	v, ok = stack.Pop()
	fmt.Println("Pop empty:", v, ok)
}
