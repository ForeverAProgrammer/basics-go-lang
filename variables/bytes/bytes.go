package main

import (
	"fmt"
	"reflect"
)

func main() {
	// byte is an alias for uint8, representing a single byte (0-255)
	var b1 byte = 65  // numeric value (ASCII for 'A')
	var b2 byte = 'B' // character literal

	fmt.Println("b1:", b1, "| type:", reflect.TypeOf(b1))
	fmt.Println("b2:", b2, "| type:", reflect.TypeOf(b2))
	fmt.Printf("b1 as character: %c\n", b1)
	fmt.Printf("b2 as character: %c\n", b2)
	fmt.Println()

	// Byte slices are commonly used to work with raw data or strings
	data := []byte{72, 101, 108, 108, 111} // ASCII for "Hello"
	fmt.Println("Byte slice:", data, "| type:", reflect.TypeOf(data))
	fmt.Println("As string:", string(data))
}
