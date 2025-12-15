# Comparison Operators Demo

This program demonstrates comparison operators in Go.

## What it does

The program shows how to use:

1. **Equal to (==)** - Checks if two values are equal
2. **Not equal to (!=)** - Checks if two values are different
3. **Greater than (>)** - Checks if left value is greater than right value
4. **Less than (<)** - Checks if left value is less than right value
5. **Greater than or equal to (>=)** - Checks if left value is greater than or equal to right value
6. **Less than or equal to (<=)** - Checks if left value is less than or equal to right value
7. **String comparisons** - Comparing strings alphabetically
8. **Boolean comparisons** - Comparing boolean values

## How to run

### Run directly (without creating an executable):
```bash
go run comparison-operators-demo.go
```

### Compile and run:
```bash
go build comparison-operators-demo.go
./comparison-operators-demo
```

## Expected output

```
Comparison Operators Demo
Using a = 10 , b = 5 , c = 10

Equal to (==):
10 == 5 is false
10 == 10 is true

Not equal to (!=):
10 != 5 is true
10 != 10 is false

Greater than (>):
10 > 5 is true
5 > 10 is false

Less than (<):
10 < 5 is false
5 < 10 is true

Greater than or equal to (>=):
10 >= 5 is true
10 >= 10 is true
5 >= 10 is false

Less than or equal to (<=):
5 <= 10 is true
10 <= 10 is true
10 <= 5 is false

Comparing Strings:
apple == banana is false
apple == apple is true
apple < banana is true

Comparing Booleans:
true == false is false
true == true is true
true != false is true
```

## Key concepts

- **Comparison result**: All comparison operators return a boolean value (`true` or `false`)
- **String comparison**: Strings are compared lexicographically (alphabetically)
- **Type safety**: Go requires both operands to be of the same type when comparing
- **Common use**: Comparison operators are typically used in `if` statements and loops to control program flow
- **Double equals**: Use `==` for comparison (not `=` which is for assignment)
