// Demonstrates error creation, wrapping, Is/As, custom errors, and recover.
package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrNotFound = errors.New("not found")

type ParseError struct {
	Line int
	Msg  string
	Err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at line %d: %s", e.Line, e.Msg)
}

func (e *ParseError) Unwrap() error { return e.Err }

func readConfig(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("read config: %w", ErrNotFound)
	}
	return "config data", nil
}

func load(path string) error {
	data, err := readConfig(path)
	if err != nil {
		return fmt.Errorf("load config %q: %w", path, err)
	}
	fmt.Println("loaded:", data)
	return nil
}

func mayPanic() {
	panic("something broke")
}

func safe() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	mayPanic()
	return nil
}

func main() {
	if err := load(""); err != nil {
		fmt.Println("error:", err)
		fmt.Println("is ErrNotFound:", errors.Is(err, ErrNotFound))
	}

	// errors.As with a wrapped custom error
	inner := errors.New("bad token")
	wrapped := fmt.Errorf("outer: %w", &ParseError{Line: 3, Msg: "bad syntax", Err: inner})
	var pe *ParseError
	if errors.As(wrapped, &pe) {
		fmt.Println("as ParseError: line", pe.Line)
	}
	fmt.Println("is inner:", errors.Is(wrapped, inner))

	// errors.Join
	joined := errors.Join(errors.New("first"), errors.New("second"))
	fmt.Println("joined:", joined)

	// os errors
	_, err := os.Open("/does/not/exist")
	fmt.Println("is not exist:", errors.Is(err, os.ErrNotExist))

	// panic -> recover -> error
	fmt.Println("safe():", safe())
}
