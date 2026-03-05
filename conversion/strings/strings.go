package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// --- int to string — common gotcha ---
	// string(n) does NOT produce the number as text.
	// It returns the UTF-8 character for that Unicode code point.
	// Use strconv.Itoa() or fmt.Sprintf() to get the numeric string.
	n := 65
	gotcha := string(n)
	correct := strconv.Itoa(n)
	sprintf := fmt.Sprintf("%d", n)

	fmt.Println("int to string:")
	fmt.Printf("  string(65):        %q  <- Unicode code point 65 = 'A', probably not what you want\n", gotcha)
	fmt.Printf("  strconv.Itoa(65):  %q | type: %s\n", correct, reflect.TypeOf(correct))
	fmt.Printf("  fmt.Sprintf(65):   %q | type: %s\n", sprintf, reflect.TypeOf(sprintf))
	fmt.Println()

	// --- string to int ---
	// strconv.Atoi returns both the value and an error.
	// Always handle the error — the string may not be a valid number.
	s := "42"
	parsed, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("string to int:")
		fmt.Printf("  s:      %q | type: %s\n", s, reflect.TypeOf(s))
		fmt.Printf("  parsed: %d | type: %s\n", parsed, reflect.TypeOf(parsed))
	}
	fmt.Println()

	// --- invalid string to int ---
	// On failure, the int value is 0 (zero value) and err is non-nil.
	bad, err := strconv.Atoi("hello")
	fmt.Println("Invalid string to int:")
	if err != nil {
		fmt.Printf("  err: %s\n", err)
		fmt.Printf("  bad: %d  <- zero value returned on error\n", bad)
	}
	fmt.Println()

	// --- float64 to string ---
	f := 3.14159
	fStr := strconv.FormatFloat(f, 'f', 2, 64) // 'f' = decimal, 2 = decimal places, 64 = float64

	fmt.Println("float64 to string:")
	fmt.Printf("  f:    %f | type: %s\n", f, reflect.TypeOf(f))
	fmt.Printf("  fStr: %q | type: %s\n", fStr, reflect.TypeOf(fStr))
	fmt.Println()

	// --- string to float64 ---
	fParsed, err := strconv.ParseFloat("3.14159", 64) // 64 = parse as float64
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("string to float64:")
		fmt.Printf("  fParsed: %f | type: %s\n", fParsed, reflect.TypeOf(fParsed))
	}
}
