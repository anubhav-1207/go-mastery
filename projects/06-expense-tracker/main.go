package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".go-mastery-expenses.json")
	ledger, err := NewLedger(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "usage: expense add <amount> <category> [description]")
			os.Exit(1)
		}
		amount, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid amount")
			os.Exit(1)
		}
		description := ""
		if len(os.Args) > 4 {
			description = os.Args[4]
		}
		e, err := ledger.Add(amount, os.Args[3], description)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Println("added expense", e.ID)

	case "list":
		for _, e := range ledger.All() {
			fmt.Printf("%d: %.2f [%s] %s (%s)\n",
				e.ID, e.Amount, e.Category, e.Description,
				e.Date.Format("2006-01-02"))
		}

	case "summary":
		for _, category := range ledger.SortedCategories() {
			fmt.Printf("%-15s %.2f\n", category, ledger.TotalByCategory()[category])
		}

	case "rm":
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, "usage: expense rm <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid id")
			os.Exit(1)
		}
		if err := ledger.Remove(id); err != nil {
			if errors.Is(err, ErrNotFound) {
				fmt.Fprintln(os.Stderr, "not found")
			} else {
				fmt.Fprintln(os.Stderr, "error:", err)
			}
			os.Exit(1)
		}
		fmt.Println("removed")

	default:
		usage()
		return
	}

	if err := ledger.Save(); err != nil {
		fmt.Fprintln(os.Stderr, "save error:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: expense <add|list|summary|rm> [args]")
}
