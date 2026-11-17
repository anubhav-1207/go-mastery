package main

import (
	"errors"
	"path/filepath"
	"testing"
)

func newTestLedger(t *testing.T) *Ledger {
	t.Helper()
	l, err := NewLedger(filepath.Join(t.TempDir(), "expenses.json"))
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestAddValidation(t *testing.T) {
	l := newTestLedger(t)
	if _, err := l.Add(0, "food", ""); err == nil {
		t.Error("expected error for zero amount")
	}
	if _, err := l.Add(10, "", ""); err == nil {
		t.Error("expected error for empty category")
	}
	e, err := l.Add(10, "food", "lunch")
	if err != nil || e.ID != 0 {
		t.Fatalf("Add = %+v, %v", e, err)
	}
}

func TestSummary(t *testing.T) {
	l := newTestLedger(t)
	l.Add(10, "food", "lunch")
	l.Add(5, "food", "coffee")
	l.Add(20, "transport", "taxi")

	totals := l.TotalByCategory()
	if totals["food"] != 15 || totals["transport"] != 20 {
		t.Errorf("unexpected totals: %v", totals)
	}
	categories := l.SortedCategories()
	if len(categories) != 2 || categories[0] != "food" || categories[1] != "transport" {
		t.Errorf("unexpected categories: %v", categories)
	}
}

func TestPersistenceAndRemove(t *testing.T) {
	path := filepath.Join(t.TempDir(), "expenses.json")
	l, _ := NewLedger(path)
	l.Add(12.5, "books", "Go book")
	if err := l.Save(); err != nil {
		t.Fatal(err)
	}

	loaded, err := NewLedger(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.All()) != 1 {
		t.Fatalf("want 1 expense, got %d", len(loaded.All()))
	}
	if err := loaded.Remove(0); err != nil {
		t.Fatal(err)
	}
	if err := loaded.Remove(0); !errors.Is(err, ErrNotFound) {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}
