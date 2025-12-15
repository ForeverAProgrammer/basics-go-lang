// Package name
package main

// Importing math and strings packages
import (
	"math"
	"strings"
	"fmt"
)

// main function - entry point of the program
func main() {

	// Math.Floor returns the greatest integer value less than or equal to 2.75
	fmt.Println("Math Floor of 2.75 is:")
	fmt.Println(math.Floor(2.75)) // Returns 2

	fmt.Println() // Blank line for better readability

	// Changes the first letter of the string to uppercase (proper case)
	fmt.Println("Proper case of 'hello go' is:")
	fmt.Println(strings.Title("hello go")) // Returns "Hello Go"
}
