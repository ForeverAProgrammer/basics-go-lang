# Numeric Conversion Demo

This program demonstrates explicit numeric type conversion in Go.

## What it does

1. **No automatic conversion** - Shows the compile-time error you get when mixing types without converting
2. **int to float64** - Widening a whole number into a floating-point type
3. **float64 to int** - Truncation of the fractional part (not rounding)
4. **Negative float64 to int** - Truncation toward zero, not toward negative infinity
5. **int16 to int32** - Safe widening conversion, value is always preserved
6. **int32 to int8** - Narrowing conversion that overflows silently

## How to run

### Run directly (without creating an executable)

```bash
go run numeric.go
```

### Compile and run

```bash
go build numeric.go
./numeric
```

## Expected output

```
int to float64:
  i: 42 | type: int
  f: 42.000000 | type: float64

float64 to int (truncation, not rounding):
  f2: 9.990000 | type: float64
  i2: 9 | type: int  <- 0.99 dropped, not rounded up to 10

Negative float64 to int (truncation toward zero):
  f3: -9.990000 | type: float64
  i3: -9 | type: int  <- -9, not -10

int16 to int32 (widening — safe, value preserved):
  n16: 1000 | type: int16
  n32: 1000 | type: int32

int32 to int8 (narrowing — overflow, no error):
  big:   2147483647 | type: int32
  small: -1 | type: int8  <- overflowed silently
```

## No automatic conversion

Go never converts numeric types automatically. This is different from languages like Python and JavaScript where numeric types mix freely:

```python
# Python — automatic conversion, no error
x = 42
y = x + 2.5   # int + float → float, works silently
```

In Go, the same code is a compile error:

```go
x := 42
y := x + 2.5  // invalid operation: x + 2.5 (mismatched types int and float64)
y := float64(x) + 2.5  // correct — explicit conversion required
```

## Truncation vs rounding

`float64` can store fractional values — `int` cannot. When you convert a `float64` to `int`, Go **truncates** (drops) the fractional part rather than rounding:

| Value  | `int()` result | Note                           |
| ------ | -------------- | ------------------------------ |
| `9.1`  | `9`            | truncated                      |
| `9.9`  | `9`            | truncated, not rounded to 10   |
| `-9.9` | `-9`           | truncated toward zero, not -10 |

If you need rounding, use `math.Round()` before converting:

```go
import "math"
rounded := int(math.Round(9.9))  // 10
```

## Key concepts

- **Explicit conversion syntax**: `targetType(value)` — e.g. `float64(x)`, `int(f)`
- **Truncation**: float64 → int drops the fractional part, always toward zero
- **Widening** (small → large type): always safe, value is preserved
- **Narrowing** (large → small type): can overflow silently — Go does not panic or error
- **No implicit conversion**: every type change must be written explicitly in code
