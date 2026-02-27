# Variable Declaration Demo

This program demonstrates the different ways to declare variables in Go.

## What it does

1. **Zero values** - Variables declared without a value are automatically initialised to their type's zero value
2. **Single-line `var` (explicit type)** - `var age int = 30`
3. **Single-line `var` (inferred type)** - `var score = 100`, type is inferred from the value
4. **Grouped `var` block** - Idiomatic way to declare several related variables together
5. **Short assignment `:=`** - Most common form inside functions, type always inferred
6. **Multi-variable short assignment** - Declare multiple variables of different types in one line
7. **Wrong type assignment** - Commented examples showing compile-time errors for type mismatches

## How to run

### Run directly (without creating an executable):
```bash
go run declaration.go
```

### Compile and run:
```bash
go build declaration.go
./declaration
```

## Expected output

```
Zero Values:
int: 0 | type: int
float64: 0.000000 | type: float64
string: '' | type: string
bool: false | type: bool

Single-line var (explicit type):
age: 30 | type: int
username: 'Alice' | type: string

Single-line var (inferred type):
score: 100 | type: int
label: 'beginner' | type: string

Grouped var block:
city: 'London' | type: string
country: 'UK' | type: string
pop: 9000000 | type: int

Short assignment :=
x: 42 | type: int
y: 3.140000 | type: float64
name: 'Go' | type: string
isActive: true | type: bool

Multi-variable short assignment:
lang: 'Go' | type: string
year: 2024 | type: int
active: false | type: bool
```

## Wrong type assignment — compile-time errors

Go is statically typed, which means type mismatches are caught **at compile time** before the program ever runs. There is no runtime exception to catch — the program simply won't build.

Uncommenting any of these lines in the source will produce a compile error:

```go
var n int = "42"   // cannot use "42" (untyped string constant) as int value
var n int = 3.14   // cannot use 3.14 (untyped float constant) as int value
x = "hello"        // cannot use "hello" (untyped string constant) as int value in assignment
```

This is different from dynamically typed languages like Python or JavaScript, where this kind of error only surfaces at runtime:

```python
# Python — no error until this line actually executes
x = 42
x = "hello"  # valid, just reassigns x to a different type
```

In Go, the compiler enforces types across the entire codebase before a single line runs.

## Naming rules

### Exported vs unexported

Whether a name starts with an uppercase or lowercase letter controls its visibility outside the package.

**Exported** — starts with a capital letter, accessible from any other package:

```go
var MaxRetries int = 3
var HTTPClient = ...
```

**Unexported** — starts with a lowercase letter, only accessible within the same package:

```go
var maxRetries int = 3
var httpClient = ...
```

This applies to variables, functions, types, and struct fields — not just variables.

### Community conventions

The Go community follows a consistent style documented in [Effective Go](https://go.dev/doc/effective_go#names):

| Rule | Do | Don't |
| --- | --- | --- |
| Use camelCase | `firstName` | `first_name` |
| Acronyms stay uppercase | `userID`, `HTTPServer` | `userId`, `HttpServer` |
| Short names for short scopes | `i`, `n`, `err`, `ok` | `index`, `number`, `error` |
| Booleans read naturally | `isActive`, `hasError` | `active`, `error` |
| Avoid package name stutter | `user.Name` | `user.UserName` |
| Exported errors prefixed with `Err` | `ErrNotFound` | `NotFoundError` |

## Key concepts

- **Zero values**: `int` → `0`, `float64` → `0.0`, `string` → `""`, `bool` → `false`
- **`var` vs `:=`**: Both declare variables; `var` works at package level and inside functions, `:=` only works inside functions
- **Type inference**: Go can infer the type from the assigned value — you don't always need to write it explicitly
- **Grouped `var` block**: Use `var (...)` to group related declarations — cleaner than repeating `var` on every line
- **Static typing**: Every variable has a fixed type set at compile time — mismatches are errors, not warnings
