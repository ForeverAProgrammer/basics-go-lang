# String Conversion Demo

This program demonstrates converting between strings and numeric types in Go.

## What it does

1. **int to string gotcha** - `string(n)` gives a Unicode character, not the number as text
2. **int to string (correct)** - Using `strconv.Itoa()` and `fmt.Sprintf()`
3. **string to int** - Using `strconv.Atoi()` with error handling
4. **Invalid string to int** - Shows what happens when the string is not a valid number
5. **float64 to string** - Using `strconv.FormatFloat()`
6. **string to float64** - Using `strconv.ParseFloat()`

## How to run

### Run directly (without creating an executable)

```bash
go run strings.go
```

### Compile and run

```bash
go build strings.go
./strings
```

## Expected output

```
int to string:
  string(65):        "A"  <- Unicode code point 65 = 'A', probably not what you want
  strconv.Itoa(65):  "65" | type: string
  fmt.Sprintf(65):   "65" | type: string

string to int:
  s:      "42" | type: string
  parsed: 42 | type: int

Invalid string to int:
  err: strconv.Atoi: parsing "hello": invalid syntax
  bad: 0  <- zero value returned on error

float64 to string:
  f:    3.141590 | type: float64
  fStr: "3.14" | type: string

string to float64:
  fParsed: 3.141590 | type: float64
```

## The `string(n)` gotcha

Using `string()` on an integer does **not** produce the number as text. It produces the UTF-8 character for that Unicode code point — the same as `rune` to string conversion:

```go
string(65)   // "A"  — Unicode code point 65 is the letter A
string(9786) // "☺"  — Unicode code point 9786 is a smiley face
```

To convert a number to its text representation, use `strconv` or `fmt`:

```go
strconv.Itoa(65)       // "65"
fmt.Sprintf("%d", 65)  // "65"
```

## Error handling with `strconv`

`strconv.Atoi` and `strconv.ParseFloat` return two values — the result and an error. The error is non-nil when the string cannot be parsed. On failure, the numeric value returned is `0` (the zero value):

```go
n, err := strconv.Atoi("hello")
// err: strconv.Atoi: parsing "hello": invalid syntax
// n:   0
```

Always check the error before using the result.

## Key concepts

- **`strconv.Itoa(n)`**: int → string (Itoa = "Integer to ASCII")
- **`strconv.Atoi(s)`**: string → int (Atoi = "ASCII to Integer"), returns `(int, error)`
- **`strconv.FormatFloat(f, 'f', precision, 64)`**: float64 → string
- **`strconv.ParseFloat(s, 64)`**: string → float64, returns `(float64, error)`
- **`fmt.Sprintf("%d", n)`**: alternative for int → string using a format verb
- **`string(n)`**: converts an int to its Unicode character — not its numeric text
