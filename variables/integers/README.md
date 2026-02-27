# Integers Demo

This program demonstrates integer types in Go.

## What it does

1. **Signed integers** - `int`, `int8`, `int16`, `int32`, `int64`
2. **Unsigned integers** - `uint`, `uint8`, `uint16`, `uint32`
3. **Inspect types** - Using `reflect.TypeOf()` to confirm each type at runtime
4. **Type bounds** - Using the `math` package constants to show the min/max value of each type

## How to run

### Run directly (without creating an executable):
```bash
go run integers.go
```

### Compile and run:
```bash
go build integers.go
./integers
```

## Expected output

```
int: 42 | type: int
int8: 127 | type: int8
int16: 32767 | type: int16
int32: 2147483647 | type: int32
int64: 9223372036854775807 | type: int64

uint: 42 | type: uint
uint8: 255 | type: uint8
uint16: 65535 | type: uint16
uint32: 4294967295 | type: uint32

Max int8: 127
Min int8: -128
Max int16: 32767
Max int32: 2147483647
Max int64: 9223372036854775807
Max uint8: 255
```

## Go is statically typed

Go is a **statically typed** language, which means every variable has a fixed type that is set at compile time and cannot change at runtime.

This is different from dynamically typed languages like Python or JavaScript, where a variable can hold any type of value and change type freely:

```python
# Python — dynamically typed
x = 42       # int
x = "hello"  # now a string — perfectly valid
```

In Go, that is a compile error:

```go
// Go — statically typed
var x int = 42
x = "hello" // compile error: cannot use "hello" (type string) as type int
```

**Why it matters for integers specifically:**

Go doesn't have one "number" type — it has many (`int8`, `int16`, `int32`, `int64`, `uint8`, etc.), and they are all distinct types. You can't mix them without an explicit conversion:

```go
var a int32 = 100
var b int64 = 200
// b = a         // compile error: cannot use a (type int32) as type int64
b = int64(a)     // correct: explicit conversion required
```

This catches overflow bugs and unintended data loss at compile time rather than silently at runtime.

## Key concepts

- **`int`**: Size depends on the platform — 32-bit on 32-bit systems, 64-bit on 64-bit systems
- **Signed vs unsigned**: Signed types (`int*`) hold negatives; unsigned types (`uint*`) sacrifice the negative range to double the positive range
- **`reflect.TypeOf()`**: Confirms the exact type — useful for seeing that `int32` and `int64` are genuinely different types
- **`math` constants**: `math.MaxInt8`, `math.MinInt8`, etc. give the exact bounds for each type without having to memorise them
