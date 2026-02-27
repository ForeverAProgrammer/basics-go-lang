package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Signed integers
	var i int = 42                     // platform-dependent (32 or 64 bit)
	var i8 int8 = 127                  // -128 to 127
	var i16 int16 = 32767              // -32768 to 32767
	var i32 int32 = 2147483647         // -2147483648 to 2147483647
	var i64 int64 = 9223372036854775807 // largest signed 64-bit value

	fmt.Println("int:", i, "| type:", reflect.TypeOf(i))
	fmt.Println("int8:", i8, "| type:", reflect.TypeOf(i8))
	fmt.Println("int16:", i16, "| type:", reflect.TypeOf(i16))
	fmt.Println("int32:", i32, "| type:", reflect.TypeOf(i32))
	fmt.Println("int64:", i64, "| type:", reflect.TypeOf(i64))
	fmt.Println()

	// Unsigned integers (no negatives, double the positive range)
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295

	fmt.Println("uint:", u, "| type:", reflect.TypeOf(u))
	fmt.Println("uint8:", u8, "| type:", reflect.TypeOf(u8))
	fmt.Println("uint16:", u16, "| type:", reflect.TypeOf(u16))
	fmt.Println("uint32:", u32, "| type:", reflect.TypeOf(u32))
	fmt.Println()

	// math package constants for max/min values
	fmt.Println("Max int8:", math.MaxInt8)
	fmt.Println("Min int8:", math.MinInt8)
	fmt.Println("Max int16:", math.MaxInt16)
	fmt.Println("Max int32:", math.MaxInt32)
	fmt.Println("Max int64:", math.MaxInt64)
	fmt.Println("Max uint8:", math.MaxUint8)
}
