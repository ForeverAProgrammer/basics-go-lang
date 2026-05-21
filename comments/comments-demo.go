package main

import "fmt"

// comments-demo is a tiny program that prints a message.
// It exists to demonstrate Go comment styles and best practices.
//
// In Go, comments fall into two main categories:
//   - Documentation comments (beginning with the name being described)
//   - Inline comments (used sparingly to explain non-obvious code)
func main() {
	// This is a single-line comment describing the next statement.
	// Use inline comments like this only when the intent is not obvious
	// from the code or when you need to call out important behavior.
	fmt.Println(Greet("world"))

	// TODO: add more comment style examples as the demo grows

	// NOTE: block comments are also valid Go syntax (shown below).
	// They are rarely used inside functions — prefer // for inline comments.
	/* This is a block comment. It can span multiple lines.
	   Block comments are most common as package-level doc comments
	   in older Go code, but the // style is preferred today. */
}

// Greet returns a greeting string for the given name.
//
// This is a documentation comment. It starts with the function name and is
// picked up by go doc and godoc. Run `go doc Greet` in this directory to see it.
func Greet(name string) string {
	return "Hello, " + name + "!"
}
