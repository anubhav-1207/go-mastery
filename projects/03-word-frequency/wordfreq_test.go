package main

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count("Go is great. GO is fun!")
	want := map[string]int{"go": 2, "is": 2, "great": 1, "fun": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Count = %v; want %v", got, want)
	}
}

func TestTop(t *testing.T) {
	freq := map[string]int{"a": 1, "b": 3, "c": 3, "d": 2}
	got := Top(freq, 3)
	want := []WordCount{{"b", 3}, {"c", 3}, {"d", 2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Top = %v; want %v", got, want)
	}
}

func TestCountEmpty(t *testing.T) {
	if got := Count(""); len(got) != 0 {
		t.Errorf("expected empty map, got %v", got)
	}
}
