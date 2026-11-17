package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// ErrNotFound is returned when a task id does not exist.
var ErrNotFound = errors.New("task not found")

// Task is a single todo item.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Store persists tasks to a JSON file.
type Store struct {
	path  string
	tasks []Task
	next  int
}

// NewStore loads tasks from path (or starts empty if the file does not exist).
func NewStore(path string) (*Store, error) {
	s := &Store{path: path, tasks: []Task{}}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, fmt.Errorf("read store: %w", err)
	}
	if err := json.Unmarshal(data, &s.tasks); err != nil {
		return nil, fmt.Errorf("parse store: %w", err)
	}
	for _, t := range s.tasks {
		if t.ID >= s.next {
			s.next = t.ID + 1
		}
	}
	return s, nil
}

// Add creates a new task.
func (s *Store) Add(title string) Task {
	t := Task{ID: s.next, Title: title, CreatedAt: time.Now()}
	s.next++
	s.tasks = append(s.tasks, t)
	return t
}

// List returns all tasks.
func (s *Store) List() []Task {
	return append([]Task(nil), s.tasks...) // return a copy
}

// MarkDone marks a task as done.
func (s *Store) MarkDone(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return ErrNotFound
}

// Remove deletes a task by id.
func (s *Store) Remove(id int) error {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// Save writes tasks back to disk.
func (s *Store) Save() error {
	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}
	if err := os.WriteFile(s.path, data, 0644); err != nil {
		return fmt.Errorf("write store: %w", err)
	}
	return nil
}
