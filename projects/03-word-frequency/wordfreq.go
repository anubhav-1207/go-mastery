package main

import (
	"sort"
	"strings"
	"unicode"
)

// Count returns a map of lowercased word -> frequency in text.
func Count(text string) map[string]int {
	freq := make(map[string]int)
	for _, word := range strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}) {
		w := strings.ToLower(word)
		if w != "" {
			freq[w]++
		}
	}
	return freq
}

// Top returns the n most frequent words, sorted by count descending then word ascending.
func Top(freq map[string]int, n int) []WordCount {
	items := make([]WordCount, 0, len(freq))
	for word, count := range freq {
		items = append(items, WordCount{Word: word, Count: count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return items[i].Word < items[j].Word
		}
		return items[i].Count > items[j].Count
	})
	if n < len(items) {
		items = items[:n]
	}
	return items
}

// WordCount is one entry in a frequency report.
type WordCount struct {
	Word  string
	Count int
}
