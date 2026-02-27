# Booleans Demo

This program demonstrates the `bool` type in Go.

## What it does

1. **Declare booleans** - Using explicit `var` declaration and short `:=` declaration
2. **Inspect types** - Using `reflect.TypeOf()` to confirm the type at runtime
3. **Boolean expressions** - AND (`&&`), OR (`||`), and NOT (`!`) operators

## How to run

### Run directly (without creating an executable):
```bash
go run booleans.go
```

### Compile and run:
```bash
go build booleans.go
./booleans
```

## Expected output

```
b1: true | type: bool
b2: false | type: bool
b3: true | type: bool

true && false: false
true || false: true
!true: false
```

## Key concepts

- **Only two values**: A `bool` can only ever be `true` or `false`
- **Type inference**: When using `:=`, Go infers the type as `bool` automatically
- **`reflect.TypeOf()`**: Returns the runtime type of any value — useful for confirming what type Go assigned
- **Logical operators**: `&&` (AND), `||` (OR), `!` (NOT) all return `bool`
- **Common use**: Booleans are used in `if` statements, loops, and flags to control program flow
