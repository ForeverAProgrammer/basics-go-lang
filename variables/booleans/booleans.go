package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Boolean type can only be true or false
	var b1 bool = true
	var b2 bool = false
	b3 := true // short declaration infers bool

	fmt.Println("b1:", b1, "| type:", reflect.TypeOf(b1))
	fmt.Println("b2:", b2, "| type:", reflect.TypeOf(b2))
	fmt.Println("b3:", b3, "| type:", reflect.TypeOf(b3))
	fmt.Println()

	// Boolean expressions
	fmt.Println("true && false:", true && false) // AND
	fmt.Println("true || false:", true || false) // OR
	fmt.Println("!true:", !true)                 // NOT
}
