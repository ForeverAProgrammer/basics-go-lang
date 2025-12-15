package main

import (
	"fmt"
)

func main() {
	// \n - New line character (moves to next line)
	fmt.Println("New Line Character Demo:")
	fmt.Println("Hello,\nGo!")

	// Print an empty line for better readability
	fmt.Println()

	// \t - Tab character (horizontal tab, typically 8 spaces)
	fmt.Println("Tab Character Demo:")
	fmt.Println("Hello,\tGo!")

	// Print an empty line for better readability
	fmt.Println()

	// Practical example: Using tabs to format output in columns
	fmt.Println("Formatting with Tabs:")
	fmt.Println("Name\tAge\tCity")
	fmt.Println("Alice\t25\tNew York")
	fmt.Println("Bob\t30\tSan Francisco")

	// Print an empty line for better readability
	fmt.Println()

	// \" - Double quote character (escaped to include in string)
	fmt.Println("Quotation Marks Demo: \"\"")

	// Print an empty line for better readability
	fmt.Println()

	// Single quote character (doesn't need escaping in double-quoted strings)
	fmt.Println("Single Quote Demo: '")

	// Print an empty line for better readability
	fmt.Println()

	// \\ - Backslash character (escaped to show literal backslash)
	fmt.Println("Backslash Demo: \\")

	// Print an empty line for better readability
	fmt.Println()

	// Raw strings with backticks - no escape processing
	fmt.Println("Raw String Demo (using backticks):")
	fmt.Println(`C:\Users\name\Documents\file.txt`)  // Backslashes don't need escaping
	fmt.Println(`Line 1
Line 2
Line 3`)  // Can span multiple lines without \n
}