package main

// Parser known how to Parse a dialogue
type Parser interface {
	// Concrete types should implement this method.
	Parse(s string) string
}
