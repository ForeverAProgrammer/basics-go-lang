package main

import (
	"fmt"
)

func main() {
	// Define numbers for comparison
	a := 10
	b := 5
	c := 10

	fmt.Println("Comparison Operators Demo")
	fmt.Println("Using a =", a, ", b =", b, ", c =", c)
	fmt.Println()

	// Equal to (==)
	fmt.Println("Equal to (==):")
	fmt.Println(a, "==", b, "is", a == b) // false
	fmt.Println(a, "==", c, "is", a == c) // true
	fmt.Println()

	// Not equal to (!=)
	fmt.Println("Not equal to (!=):")
	fmt.Println(a, "!=", b, "is", a != b) // true
	fmt.Println(a, "!=", c, "is", a != c) // false
	fmt.Println()

	// Greater than (>)
	fmt.Println("Greater than (>):")
	fmt.Println(a, ">", b, "is", a > b) // true
	fmt.Println(b, ">", a, "is", b > a) // false
	fmt.Println()

	// Less than (<)
	fmt.Println("Less than (<):")
	fmt.Println(a, "<", b, "is", a < b) // false
	fmt.Println(b, "<", a, "is", b < a) // true
	fmt.Println()

	// Greater than or equal to (>=)
	fmt.Println("Greater than or equal to (>=):")
	fmt.Println(a, ">=", b, "is", a >= b) // true
	fmt.Println(a, ">=", c, "is", a >= c) // true (equal values)
	fmt.Println(b, ">=", a, "is", b >= a) // false
	fmt.Println()

	// Less than or equal to (<=)
	fmt.Println("Less than or equal to (<=):")
	fmt.Println(b, "<=", a, "is", b <= a) // true
	fmt.Println(a, "<=", c, "is", a <= c) // true (equal values)
	fmt.Println(a, "<=", b, "is", a <= b) // false
	fmt.Println()

	// Comparing strings
	// Strings are compared lexicographically (alphabetically) character by character
	fmt.Println("Comparing Strings:")
	str1 := "apple"
	str2 := "banana"
	str3 := "apple"

	fmt.Println(str1, "==", str2, "is", str1 == str2) // false
	fmt.Println(str1, "==", str3, "is", str1 == str3) // true
	fmt.Println(str1, "<", str2, "is", str1 < str2)   // true ('a' comes before 'b' alphabetically)
	fmt.Println()

	// Comparing booleans
	fmt.Println("Comparing Booleans:")
	bool1 := true
	bool2 := false
	bool3 := true

	fmt.Println(bool1, "==", bool2, "is", bool1 == bool2) // false
	fmt.Println(bool1, "==", bool3, "is", bool1 == bool3) // true
	fmt.Println(bool1, "!=", bool2, "is", bool1 != bool2) // true
}
