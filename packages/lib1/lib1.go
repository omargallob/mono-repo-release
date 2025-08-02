package lib1

import "fmt"

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello %s from lib1 (version %s)", name, Version)
}

// Add calculates the sum of two integers
func Add(a, b int) int {
	return a + b
}
