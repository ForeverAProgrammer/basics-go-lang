package main

import (
	"fmt"
)

func main() {
	// Define two numbers for our operations
	a := 10
	b := 3

	fmt.Println("Math Operations Demo")
	fmt.Println("Using a =", a, "and b =", b)
	fmt.Println()

	// Addition (+)
	fmt.Println("Addition:")
	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println()

	// Subtraction (-)
	fmt.Println("Subtraction:")
	fmt.Println(a, "-", b, "=", a-b)
	fmt.Println()

	// Multiplication (*)
	fmt.Println("Multiplication:")
	fmt.Println(a, "*", b, "=", a*b)
	fmt.Println()

	// Division (/)
	fmt.Println("Division:")
	fmt.Println(a, "/", b, "=", a/b) // Integer division (result is 3, not 3.33...)
	fmt.Println()

	// Division with floats for decimal result
	fmt.Println("Division with floats:")
	fmt.Println(float64(a), "/", float64(b), "=", float64(a)/float64(b))
	fmt.Println()

	// Modulus (%) - remainder after division
	fmt.Println("Modulus (Remainder):")
	fmt.Println(a, "%", b, "=", a%b) // 10 divided by 3 leaves remainder 1
	fmt.Println()

	// Increment and Decrement
	fmt.Println("Increment and Decrement:")
	x := 5
	fmt.Println("x starts at:", x)
	x++ // Increment by 1
	fmt.Println("After x++:", x)
	x-- // Decrement by 1
	fmt.Println("After x--:", x)
	fmt.Println()

	// Compound assignment operators
	fmt.Println("Compound Assignment Operators:")
	y := 10
	fmt.Println("y starts at:", y)

	y += 5 // Same as y = y + 5
	fmt.Println("After y += 5:", y)

	y -= 3 // Same as y = y - 3
	fmt.Println("After y -= 3:", y)

	y *= 2 // Same as y = y * 2
	fmt.Println("After y *= 2:", y)

	y /= 4 // Same as y = y / 4
	fmt.Println("After y /= 4:", y)

	y %= 3 // Same as y = y % 3
	fmt.Println("After y %= 3:", y)
}
