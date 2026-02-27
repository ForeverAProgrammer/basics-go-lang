# Floats Demo

This program demonstrates floating-point types in Go.

## What it does

1. **Declare floats** - `float32` and `float64` with the same literal value
2. **Inspect types** - Using `reflect.TypeOf()` to confirm each type at runtime
3. **Precision difference** - Shows how `float32` loses precision vs `float64`
4. **Type bounds** - Using `math` package constants for max and smallest nonzero values

## How to run

### Run directly (without creating an executable):
```bash
go run floats.go
```

### Compile and run:
```bash
go build floats.go
./floats
```

## Expected output

```
float32: 3.1415927 | type: float32
float64: 3.14159265358979 | type: float64

pi (float64): 3.141592653589793 | type: float64

Max float32: 3.4028234663852886e+38
Max float64: 1.7976931348623157e+308
Smallest nonzero float32: 1.401298464324817e-45
Smallest nonzero float64: 5e-324
```

## Key concepts

- **Default type**: Floating-point literals (e.g. `3.14`) default to `float64` when using `:=`
- **`float32`**: ~6–7 significant decimal digits of precision
- **`float64`**: ~15–17 significant decimal digits of precision — preferred in most cases
- **Precision loss**: Both variables are assigned the same literal, but `float32` rounds it — visible in the output
- **`reflect.TypeOf()`**: Confirms which float type was assigned, especially useful with `:=` where the type is inferred
- **`math` constants**: `math.MaxFloat32`, `math.MaxFloat64`, `math.SmallestNonzeroFloat32`, `math.SmallestNonzeroFloat64`
