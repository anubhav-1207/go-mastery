// Package mathutil provides simple arithmetic helpers for the tutorial.
package mathutil

// Pi is a useful constant. Exported because it starts with an uppercase letter.
const Pi = 3.14159

// version is package-private: it cannot be seen outside this package.
const version = "1.0.0"

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Rectangle holds dimensions.
type Rectangle struct {
	Width, Height float64
}

// Area returns the rectangle's area.
func (r Rectangle) Area() float64 { return r.Width * r.Height }

// Version returns the package version string.
func Version() string { return version }
