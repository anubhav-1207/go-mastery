// Package perf contains benchmarks demonstrating allocation-aware code.
package perf

import (
	"strings"
	"testing"
)

var sink any

func BenchmarkAppendNoPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var out []int
		for j := 0; j < 1000; j++ {
			out = append(out, j)
		}
		sink = out
	}
}

func BenchmarkAppendPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := make([]int, 0, 1000)
		for j := 0; j < 1000; j++ {
			out = append(out, j)
		}
		sink = out
	}
}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "x"
		}
		sink = s
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.Grow(100)
		for j := 0; j < 100; j++ {
			sb.WriteString("x")
		}
		sink = sb.String()
	}
}

func TestAllocsPerRun(t *testing.T) {
	// Guard against accidental allocations in a small helper.
	allocs := testing.AllocsPerRun(100, func() { sink = add(1, 2) })
	if allocs != 0 {
		t.Errorf("add allocated %v times per run", allocs)
	}
}

func add(a, b int) int { return a + b }
