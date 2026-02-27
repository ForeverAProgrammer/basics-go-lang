package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Strings are immutable sequences of UTF-8 encoded bytes
	var s1 string = "Hello, World!"
	s2 := "Go is awesome" // short declaration

	fmt.Println("s1:", s1, "| type:", reflect.TypeOf(s1))
	fmt.Println("s2:", s2, "| type:", reflect.TypeOf(s2))
	fmt.Println()

	// len() returns the number of bytes, not characters
	fmt.Println("Length of s1:", len(s1))
	fmt.Println()

	// String concatenation with +
	s3 := s1 + " " + s2
	fmt.Println("Concatenated:", s3)
	fmt.Println()

	// Raw string literal using backticks — no escape sequences processed
	s4 := `This is a
multi-line
string`
	fmt.Println("Raw string:", s4)
	fmt.Println()

	// Accessing a byte at a specific index
	fmt.Println("First byte of s1:", s1[0]) // prints the numeric byte value
	fmt.Printf("First character of s1: %c\n", s1[0])
}
