package main

import (
	"fmt"
	"reflect"
)

func main() {
	// --- No automatic conversion ---
	// Unlike Python or JavaScript, Go never converts numeric types automatically.
	// Mixing types without an explicit conversion is a compile-time error:
	//
	//   var x int = 42
	//   var y float64 = x       // cannot use x (type int) as type float64
	//   result := x + 2.5       // cannot use 2.5 (untyped float) as type int
	//
	// Every conversion must be written explicitly: float64(x), int(y), etc.

	// --- int to float64 ---
	i := 42
	f := float64(i)

	fmt.Println("int to float64:")
	fmt.Printf("  i: %d | type: %s\n", i, reflect.TypeOf(i))
	fmt.Printf("  f: %f | type: %s\n", f, reflect.TypeOf(f))
	fmt.Println()

	// --- float64 to int ---
	// float64 can store fractional values — int cannot.
	// When converting float64 to int the fractional part is dropped (truncated
	// toward zero), it is NOT rounded to the nearest whole number.
	f2 := 9.99
	i2 := int(f2)

	fmt.Println("float64 to int (truncation, not rounding):")
	fmt.Printf("  f2: %f | type: %s\n", f2, reflect.TypeOf(f2))
	fmt.Printf("  i2: %d | type: %s  <- 0.99 dropped, not rounded up to 10\n", i2, reflect.TypeOf(i2))
	fmt.Println()

	// Truncation always moves toward zero, not toward negative infinity
	f3 := -9.99
	i3 := int(f3)

	fmt.Println("Negative float64 to int (truncation toward zero):")
	fmt.Printf("  f3: %f | type: %s\n", f3, reflect.TypeOf(f3))
	fmt.Printf("  i3: %d | type: %s  <- -9, not -10\n", i3, reflect.TypeOf(i3))
	fmt.Println()

	// --- int to int (widening — always safe) ---
	var n16 int16 = 1000
	n32 := int32(n16)

	fmt.Println("int16 to int32 (widening — safe, value preserved):")
	fmt.Printf("  n16: %d | type: %s\n", n16, reflect.TypeOf(n16))
	fmt.Printf("  n32: %d | type: %s\n", n32, reflect.TypeOf(n32))
	fmt.Println()

	// --- int to int (narrowing — can overflow silently) ---
	// Go does not panic or error on overflow — the bits are simply truncated.
	// It is your responsibility to ensure the value fits.
	var big int32 = 2147483647 // max value for int32
	small := int8(big)         // does not fit in int8 — overflows silently

	fmt.Println("int32 to int8 (narrowing — overflow, no error):")
	fmt.Printf("  big:   %d | type: %s\n", big, reflect.TypeOf(big))
	fmt.Printf("  small: %d | type: %s  <- overflowed silently\n", small, reflect.TypeOf(small))
}
