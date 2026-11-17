package main

import (
	"errors"
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		op      string
		a, b    float64
		want    float64
		wantErr error
	}{
		{"add", "add", 2, 3, 5, nil},
		{"subtract", "sub", 10, 4, 6, nil},
		{"multiply", "mul", 3, 4, 12, nil},
		{"divide", "div", 10, 2, 5, nil},
		{"power", "pow", 2, 8, 256, nil},
		{"divide by zero", "div", 1, 0, 0, ErrDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.op, tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v; want %v", err, tt.wantErr)
			}
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}

	if _, err := Calculate("nope", 1, 2); err == nil {
		t.Error("expected error for unknown operation")
	}
}
