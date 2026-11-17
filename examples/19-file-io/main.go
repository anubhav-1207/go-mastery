// Demonstrates reading/writing files, bufio, and filepath.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.MkdirTemp("", "go-mastery-io-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "notes.txt")
	fmt.Println("working file:", path)

	// Write a whole file
	content := strings.Repeat("line of text\n", 5)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		panic(err)
	}

	// Read a whole file
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("bytes read:", len(data))

	// Read line by line
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	lineCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineCount++
		_ = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	f.Close()
	fmt.Println("lines:", lineCount)

	// Buffered writing
	outPath := filepath.Join(dir, "out.txt")
	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(out)
	for i := 0; i < 100; i++ {
		fmt.Fprintf(w, "item %d\n", i)
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
	out.Close()
	fmt.Println("wrote bytes:", fileSize(outPath))

	// Append
	af, err := os.OpenFile(outPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	io.WriteString(af, "appended\n")
	af.Close()

	// File info
	info, _ := os.Stat(outPath)
	fmt.Println("size:", info.Size(), "isDir:", info.IsDir())

	// Directory listing
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		fmt.Println("entry:", e.Name())
	}

	// Paths
	fmt.Println("join:", filepath.Join("a", "b", "c.txt"))
	fmt.Println("ext:", filepath.Ext("archive.tar.gz"))
	fmt.Println("base:", filepath.Base("/tmp/a.txt"))
}

func fileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}
