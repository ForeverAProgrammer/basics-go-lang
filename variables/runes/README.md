# Runes Demo

This program demonstrates the `rune` type in Go.

## What it does

1. **Declare runes** - An ASCII character and a Unicode emoji
2. **Inspect types** - Using `reflect.TypeOf()` to show that `rune` is an alias for `int32`
3. **Print as character** - Using `%c` with `fmt.Printf` to display the actual character

## How to run

### Run directly (without creating an executable):
```bash
go run runes.go
```

### Compile and run:
```bash
go build runes.go
./runes
```

## Expected output

```
Rune r1: 65 | type: int32
Rune r2: 128522 | type: int32
r1 as character: A
r2 as character: 😊
```

## Key concepts

- **`rune` is an alias for `int32`**: `reflect.TypeOf()` returns `int32`, not `rune` — they are the same underlying type
- **Unicode code point**: A rune holds a single Unicode code point — `'A'` is `65`, `'😊'` is `128522`
- **Single quotes**: Rune literals use single quotes (`'A'`), string literals use double quotes (`"A"`)
- **`%c` format verb**: Prints the character represented by the rune's numeric value
- **Rune vs byte**: A `byte` (`uint8`) can only represent 0–255; a `rune` (`int32`) covers all ~1.1 million Unicode code points
