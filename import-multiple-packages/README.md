# Import Multiple Packages Demo

This program demonstrates how to import and use multiple packages in Go.

## What it does

The program shows how to:

1. **Import multiple packages** - Uses the grouped import syntax with parentheses
2. **Use the `math` package** - Demonstrates `math.Floor()` to round down a decimal number
3. **Use the `strings` package** - Demonstrates `strings.Title()` to capitalize words (Proper Case)
4. **Use the `fmt` package** - Demonstrates how to print text to the console.

## How to run

### Run directly (without creating an executable):

```bash
go run import-multi-package-demo.go
```

### Compile and run:

```bash
go build import-multi-package-demo.go
./import-multi-package-demo
```

## Expected output

```
Math Floor of 2.75 is:
2

Proper case of 'hello go' is:
Hello Go
```

## Key concepts

- **Grouped imports**: Using parentheses to import multiple packages cleanly
- **Standard library packages**: `math`, `fmt` and `strings` are part of Go's standard library
- **Package functions**: Functions like `math.Floor()`, `fmt.Println()` and `strings.Title()` are called using the package name prefix
