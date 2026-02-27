package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Runes in Go are represented as int32 and can be used to represent Unicode characters
	var r1 rune = 'A'  // ASCII character
	var r2 rune = '😊' // Unicode emoji character

	fmt.Println("Rune r1:", r1, "| type:", reflect.TypeOf(r1))
	fmt.Println("Rune r2:", r2, "| type:", reflect.TypeOf(r2))
	fmt.Printf("r1 as character: %c\n", r1)
	fmt.Printf("r2 as character: %c\n", r2)
}