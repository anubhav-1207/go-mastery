package calc

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Table-driven test
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("Divide(10,2) = %d, %v", got, err)
	}

	_, err = Divide(1, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Fatalf("expected ErrDivisionByZero, got %v", err)
	}
}

// Test helper
func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	assertEqual(t, Reverse("hello"), "olleh")
	assertEqual(t, Reverse("世界"), "界世")
	assertEqual(t, Reverse(""), "")
}

// Benchmark
var sink string

func BenchmarkReverse(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = Reverse("hello world this is a string")
	}
	sink = s
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 100; j++ {
			sb.WriteString("x")
		}
		sink = sb.String()
	}
}

// Example (also checked by go test)
func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}
