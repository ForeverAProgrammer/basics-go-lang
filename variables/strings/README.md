# Strings Demo

This program demonstrates the `string` type in Go.

## What it does

1. **Declare strings** - Using explicit `var` and short `:=` declaration
2. **Inspect types** - Using `reflect.TypeOf()` to confirm the type at runtime
3. **String length** - Using `len()` to get the byte count
4. **Concatenation** - Joining strings with `+`
5. **Raw string literals** - Using backticks for multi-line strings with no escape processing
6. **Byte access** - Indexing into a string to read individual bytes

## How to run

### Run directly (without creating an executable):
```bash
go run strings.go
```

### Compile and run:
```bash
go build strings.go
./strings
```

## Expected output

```
s1: Hello, World! | type: string
s2: Go is awesome | type: string

Length of s1: 13

Concatenated: Hello, World! Go is awesome

Raw string: This is a
multi-line
string

First byte of s1: 72
First character of s1: H
```

## Key concepts

- **Immutable**: Strings in Go cannot be modified after creation — you create new strings instead
- **UTF-8 encoded**: Go source files and strings are UTF-8 by default
- **`len()` counts bytes, not characters**: For ASCII this is the same, but a multi-byte Unicode character (like an emoji) will count as more than 1
- **Raw string literals**: Enclosed in backticks — backslashes are treated literally and the string can span multiple lines
- **Byte indexing**: `s[i]` returns the byte at position `i` as a `uint8` — use `%c` with `fmt.Printf` to print it as a character
- **`reflect.TypeOf()`**: Confirms the type is `string`
