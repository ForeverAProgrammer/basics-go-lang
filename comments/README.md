# Comments demo

This example shows how to write **good comments in Go code**.

## What the demo does

- Demonstrates single-line, block, and documentation comment styles.
- Shows a `// TODO:` and `// NOTE:` convention used in real codebases.
- Exports a `Greet` function with a proper doc comment so you can run `go doc` against it.

## Go comment best practices

- **Prefer clear code over comments**
  - First, make the code self-explanatory with good names and small functions.
  - Add comments for *why* something is done, not to repeat *what* the code already says.

- **Use documentation comments for exported symbols**
  - For packages, types, functions, methods, and variables that are exported, start the comment with the name:
    - `// Package mypkg provides ...`
    - `// Greet returns a greeting string for the given name.`
  - This is how tools like `go doc` and `godoc` generate documentation.

- **Block comments (`/* ... */`)**
  - Valid Go syntax, but the `//` style is preferred today for almost everything.
  - Still occasionally seen in older code or for temporarily commenting out large blocks.

- **`// TODO:` and `// NOTE:` conventions**
  - `// TODO:` marks work that still needs to be done.
  - `// NOTE:` calls out something non-obvious that the reader should pay attention to.
  - These are informal conventions (not enforced by the compiler), but widely used and understood.

- **Keep comments close to the code they describe**
  - Place a comment immediately above the declaration it documents.
  - For short, local clarifications, use an inline comment at the end of the line or on the line just above.

- **Explain intent, invariants, and edge cases**
  - Document assumptions, preconditions, and postconditions that are not obvious.
  - Call out tricky logic, performance trade-offs, or non-obvious choices.

- **Avoid redundant or noisy comments**
  - Don't narrate the code (e.g., `// increment i` above `i++`).
  - Remove comments that simply repeat the function or variable name.

- **Keep comments up to date**
  - Out-of-date comments are worse than no comments.
  - When changing behavior, update or delete nearby comments as part of the same change.

## How to run

From the `comments` directory:

```bash
go run comments-demo.go
```

To see the documentation comment on `Greet` rendered by `go doc`:

```bash
go doc Greet
```
