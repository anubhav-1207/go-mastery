package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "", "input file (omit to use the built-in sample)")
	topN := flag.Int("top", 10, "number of words to show")
	flag.Parse()

	text := sampleText
	if *file != "" {
		data, err := os.ReadFile(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		text = string(data)
	}

	freq := Count(text)
	fmt.Printf("unique words: %d, total words: %d\n", len(freq), total(freq))
	for i, wc := range Top(freq, *topN) {
		fmt.Printf("%2d. %-15s %d\n", i+1, wc.Word, wc.Count)
	}
}

func total(freq map[string]int) int {
	sum := 0
	for _, n := range freq {
		sum += n
	}
	return sum
}

const sampleText = `Go is a great language. Go is simple, and Go is fast.
Concurrency in Go makes Go a great choice for servers. Simple is powerful.`
