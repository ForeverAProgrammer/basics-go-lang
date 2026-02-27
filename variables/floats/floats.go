package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// float32 has ~6-7 decimal digits of precision
	var f32 float32 = 3.14159265358979

	// float64 has ~15-17 decimal digits of precision
	// floating point literals default to float64
	var f64 float64 = 3.14159265358979

	fmt.Println("float32:", f32, "| type:", reflect.TypeOf(f32))
	fmt.Println("float64:", f64, "| type:", reflect.TypeOf(f64))
	fmt.Println()

	// Notice float32 loses precision compared to float64
	pi := 3.14159265358979323846
	fmt.Println("pi (float64):", pi, "| type:", reflect.TypeOf(pi))
	fmt.Println()

	// math package constants for float bounds
	fmt.Println("Max float32:", math.MaxFloat32)
	fmt.Println("Max float64:", math.MaxFloat64)
	fmt.Println("Smallest nonzero float32:", math.SmallestNonzeroFloat32)
	fmt.Println("Smallest nonzero float64:", math.SmallestNonzeroFloat64)
}
