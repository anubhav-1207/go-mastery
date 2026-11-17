// Demonstrates several best practices together: guard clauses, error wrapping,
// narrow interfaces, functional options, and deferred cleanup.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Store is a narrow interface defined where it is used.
type Store interface {
	Get(key string) (string, error)
}

type memoryStore map[string]string

func (m memoryStore) Get(key string) (string, error) {
	v, ok := m[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return v, nil
}

type Client struct {
	store   Store
	timeout time.Duration
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.timeout = d }
}

func NewClient(store Store, opts ...Option) *Client {
	c := &Client{store: store, timeout: 5 * time.Second}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) Lookup(key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	value, err := c.store.Get(key)
	if err != nil {
		return "", fmt.Errorf("lookup %q: %w", key, err)
	}
	return value, nil
}

func writeReport(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := io.WriteString(w, line+"\n"); err != nil {
			return fmt.Errorf("write report: %w", err)
		}
	}
	return nil
}

func main() {
	store := memoryStore{"greeting": "hello"}
	client := NewClient(store, WithTimeout(time.Second))

	value, err := client.Lookup("greeting")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", value)
	}

	_, err = client.Lookup("")
	fmt.Println("empty key error:", err)
	_, err = client.Lookup("missing")
	fmt.Println("missing error:", err)

	// Deferred cleanup with error handling
	f, err := os.CreateTemp("", "report-*.txt")
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("close error:", err)
		}
		_ = os.Remove(f.Name())
	}()

	if err := writeReport(f, []string{"first", "second"}); err != nil {
		fmt.Println("report error:", err)
		return
	}
	data, _ := os.ReadFile(f.Name())
	fmt.Println("report:", strings.TrimSpace(string(data)))
}
