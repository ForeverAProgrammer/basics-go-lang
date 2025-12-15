# Escape Characters Demo

This program demonstrates how to use escape characters and raw strings in Go.

## What it does

The program shows how to:

1. **Use `\n`** - Newline character that moves to the next line
2. **Use `\t`** - Tab character for horizontal spacing
3. **Format with tabs** - Practical example of creating columnar output
4. **Use `\"`** - Escape double quotes within strings
5. **Use single quotes** - Shows that single quotes don't need escaping in double-quoted strings
6. **Use `\\`** - Escape backslashes to display literal backslash characters
7. **Use raw strings** - Demonstrates backtick strings where escape sequences are not processed

## How to run

### Run directly (without creating an executable):
```bash
go run escape-characters-demo.go
```

### Compile and run:
```bash
go build escape-characters-demo.go
./escape-characters-demo
```

## Expected output

```
New Line Character Demo:
Hello,
Go!

Tab Character Demo:
Hello,	Go!

Formatting with Tabs:
Name	Age	City
Alice	25	New York
Bob	30	San Francisco

Quotation Marks Demo: ""

Single Quote Demo: '

Backslash Demo: \

Raw String Demo (using backticks):
C:\Users\name\Documents\file.txt
Line 1
Line 2
Line 3
```

## Key concepts

- **Escape sequences**: Special character combinations starting with `\` that represent non-printable or special characters
- **Common escape characters**: `\n` (newline), `\t` (tab), `\"` (quote), `\\` (backslash)
- **Raw strings**: Enclosed in backticks `` ` ``, they treat backslashes literally and can span multiple lines
- **Practical use**: Tabs are useful for formatting output in columns, raw strings are great for file paths and multi-line text
