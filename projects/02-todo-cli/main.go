package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".go-mastery-todo.json")

	store, err := NewStore(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		usage()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		title := strings.Join(args, " ")
		if title == "" {
			fmt.Fprintln(os.Stderr, "usage: todo add <title>")
			os.Exit(1)
		}
		t := store.Add(title)
		fmt.Println("added:", t.ID)
	case "list":
		for _, t := range store.List() {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("[%s] %d: %s\n", mark, t.ID, t.Title)
		}
	case "done", "rm":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "usage: todo done|rm <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid id")
			os.Exit(1)
		}
		if cmd == "done" {
			err = store.MarkDone(id)
		} else {
			err = store.Remove(id)
		}
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				fmt.Fprintln(os.Stderr, "task not found")
			} else {
				fmt.Fprintln(os.Stderr, "error:", err)
			}
			os.Exit(1)
		}
		fmt.Println("ok")
	default:
		usage()
		return
	}

	if err := store.Save(); err != nil {
		fmt.Fprintln(os.Stderr, "save error:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: todo <add|list|done|rm> [args]")
}
