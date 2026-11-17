package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

// ErrNotFound is returned for an unknown expense id.
var ErrNotFound = errors.New("expense not found")

// Expense is a single recorded expense.
type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Ledger manages expenses with JSON persistence.
type Ledger struct {
	path     string
	expenses []Expense
	next     int
}

// NewLedger loads a ledger from path, or creates an empty one.
func NewLedger(path string) (*Ledger, error) {
	l := &Ledger{path: path, expenses: []Expense{}}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return l, nil
		}
		return nil, fmt.Errorf("read ledger: %w", err)
	}
	if err := json.Unmarshal(data, &l.expenses); err != nil {
		return nil, fmt.Errorf("parse ledger: %w", err)
	}
	for _, e := range l.expenses {
		if e.ID >= l.next {
			l.next = e.ID + 1
		}
	}
	return l, nil
}

// Add records a new expense.
func (l *Ledger) Add(amount float64, category, description string) (Expense, error) {
	if amount <= 0 {
		return Expense{}, errors.New("amount must be positive")
	}
	if category == "" {
		return Expense{}, errors.New("category required")
	}
	e := Expense{
		ID: l.next, Amount: amount, Category: category,
		Description: description, Date: time.Now(),
	}
	l.next++
	l.expenses = append(l.expenses, e)
	return e, nil
}

// Remove deletes an expense by id.
func (l *Ledger) Remove(id int) error {
	for i, e := range l.expenses {
		if e.ID == id {
			l.expenses = append(l.expenses[:i], l.expenses[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// All returns a copy of all expenses.
func (l *Ledger) All() []Expense {
	return append([]Expense(nil), l.expenses...)
}

// TotalByCategory returns the sum of amounts per category.
func (l *Ledger) TotalByCategory() map[string]float64 {
	totals := make(map[string]float64)
	for _, e := range l.expenses {
		totals[e.Category] += e.Amount
	}
	return totals
}

// SortedCategories returns category names sorted alphabetically.
func (l *Ledger) SortedCategories() []string {
	totals := l.TotalByCategory()
	categories := make([]string, 0, len(totals))
	for c := range totals {
		categories = append(categories, c)
	}
	sort.Strings(categories)
	return categories
}

// Save writes the ledger to disk.
func (l *Ledger) Save() error {
	data, err := json.MarshalIndent(l.expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal ledger: %w", err)
	}
	if err := os.WriteFile(l.path, data, 0644); err != nil {
		return fmt.Errorf("write ledger: %w", err)
	}
	return nil
}
