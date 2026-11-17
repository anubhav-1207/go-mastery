package main

import (
	"errors"
	"sync"
)

// ErrNotFound is returned when a short code is unknown.
var ErrNotFound = errors.New("short code not found")

// Store holds short-code -> URL mappings. It is safe for concurrent use.
type Store struct {
	mu      sync.RWMutex
	urls    map[string]string
	counter int
}

// NewStore creates an empty store.
func NewStore() *Store {
	return &Store{urls: make(map[string]string)}
}

// Shorten stores a URL and returns its generated code.
func (s *Store) Shorten(url string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	code := encode(s.counter)
	s.counter++
	s.urls[code] = url
	return code
}

// Resolve returns the URL for a code.
func (s *Store) Resolve(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.urls[code]
	if !ok {
		return "", ErrNotFound
	}
	return url, nil
}

// Count returns the number of stored URLs.
func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// encode maps an integer to a short base-62 code.
func encode(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if n == 0 {
		return string(alphabet[0])
	}
	out := ""
	for n > 0 {
		out = string(alphabet[n%len(alphabet)]) + out
		n /= len(alphabet)
	}
	return out
}
