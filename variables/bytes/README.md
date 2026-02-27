# Bytes Demo

This program demonstrates the `byte` type in Go.

## What it does

1. **Declare bytes** - Using a numeric value and a character literal
2. **Inspect types** - Using `reflect.TypeOf()` to show that `byte` is an alias for `uint8`
3. **Print as character** - Using `%c` with `fmt.Printf` to display the character
4. **Byte slices** - Building a `[]byte` from numeric values and converting it to a string

## How to run

### Run directly (without creating an executable):
```bash
go run bytes.go
```

### Compile and run:
```bash
go build bytes.go
./bytes
```

## Expected output

```
b1: 65 | type: uint8
b2: 66 | type: uint8
b1 as character: A
b2 as character: B

Byte slice: [72 101 108 108 111] | type: []uint8
As string: Hello
```

## Key concepts

- **`byte` is an alias for `uint8`**: `reflect.TypeOf()` returns `uint8`, not `byte` — they are the same underlying type
- **Range 0–255**: A single byte holds values from 0 to 255, matching one ASCII character
- **Character literals**: `'B'` assigns the ASCII value `66` to the byte
- **Byte slices (`[]byte`)**: The primary way to work with raw binary data or manipulate strings at the byte level
- **`string()` conversion**: A `[]byte` can be converted directly to a `string` and vice versa
- **Byte vs rune**: Use `byte` for ASCII / raw binary data; use `rune` when working with full Unicode characters
