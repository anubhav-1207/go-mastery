package main

import (
	"errors"
	"path/filepath"
	"testing"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()
	s, err := NewStore(filepath.Join(t.TempDir(), "todo.json"))
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func TestAddList(t *testing.T) {
	s := newTestStore(t)
	s.Add("write code")
	s.Add("write tests")
	tasks := s.List()
	if len(tasks) != 2 {
		t.Fatalf("want 2 tasks, got %d", len(tasks))
	}
	if tasks[0].ID != 0 || tasks[1].ID != 1 {
		t.Errorf("unexpected IDs: %+v", tasks)
	}
}

func TestMarkDone(t *testing.T) {
	s := newTestStore(t)
	s.Add("task")
	if err := s.MarkDone(0); err != nil {
		t.Fatal(err)
	}
	if !s.List()[0].Done {
		t.Error("task should be done")
	}
	if err := s.MarkDone(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}

func TestPersistence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todo.json")
	s, _ := NewStore(path)
	s.Add("persisted")
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}

	loaded, err := NewStore(path)
	if err != nil {
		t.Fatal(err)
	}
	tasks := loaded.List()
	if len(tasks) != 1 || tasks[0].Title != "persisted" {
		t.Fatalf("unexpected loaded tasks: %+v", tasks)
	}
	if loaded.Add("next").ID != 1 {
		t.Error("ID should continue after load")
	}
}

func TestRemove(t *testing.T) {
	s := newTestStore(t)
	s.Add("a")
	s.Add("b")
	if err := s.Remove(0); err != nil {
		t.Fatal(err)
	}
	tasks := s.List()
	if len(tasks) != 1 || tasks[0].Title != "b" {
		t.Fatalf("unexpected tasks: %+v", tasks)
	}
}
