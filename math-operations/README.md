# Math Operations Demo

This program demonstrates basic mathematical operations in Go.

## What it does

The program shows how to use:

1. **Addition (+)** - Adding two numbers together
2. **Subtraction (-)** - Subtracting one number from another
3. **Multiplication (*)** - Multiplying two numbers
4. **Division (/)** - Dividing numbers (integer and float division)
5. **Modulus (%)** - Finding the remainder after division
6. **Increment (++)** - Increasing a value by 1
7. **Decrement (--)** - Decreasing a value by 1
8. **Compound assignment operators** - `+=`, `-=`, `*=`, `/=`, `%=`

## How to run

### Run directly (without creating an executable):
```bash
go run math-operations-demo.go
```

### Compile and run:
```bash
go build math-operations-demo.go
./math-operations-demo
```

## Expected output

```
Math Operations Demo
Using a = 10 and b = 3

Addition:
10 + 3 = 13

Subtraction:
10 - 3 = 7

Multiplication:
10 * 3 = 30

Division:
10 / 3 = 3

Division with floats:
10 / 3 = 3.3333333333333335

Modulus (Remainder):
10 % 3 = 1

Increment and Decrement:
x starts at: 5
After x++: 6
After x--: 5

Compound Assignment Operators:
y starts at: 10
After y += 5: 15
After y -= 3: 12
After y *= 2: 24
After y /= 4: 6
After y %= 3: 0
```

## Key concepts

- **Integer division**: When dividing two integers, the result is truncated (no decimal part)
- **Float division**: Convert to `float64` for decimal results
- **Modulus operator**: Returns the remainder after division (useful for checking even/odd numbers)
- **Increment/Decrement**: `++` and `--` are statements in Go, not expressions (must be on their own line)
- **Compound operators**: Shortcuts for common operations like `x = x + 5` becomes `x += 5`
