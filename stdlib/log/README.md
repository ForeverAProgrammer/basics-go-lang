## Log package demo

This example shows how to use Go's built-in `log` package for basic logging.

### What it covers

- **Default logger**
  - Using `log.Println` and `log.Printf`.
  - Understanding that the default logger writes to **stderr** with date/time by default.
- **Custom prefixes and flags**
  - Changing the prefix with `log.SetPrefix`.
  - Changing output format with `log.SetFlags` (e.g. `log.LstdFlags`, `log.Lshortfile`).
- **Different "log types" via prefixes**
  - Creating multiple `*log.Logger` instances with `log.New`, such as:
    - `INFO` logger writing to stdout.
    - `WARN` logger writing to stdout.
    - `ERROR` logger writing to stderr.
- **`log.Fatal` and `log.Panic`**
  - `log.Fatal` logs the message and then calls `os.Exit(1)` — the program stops immediately (no `defer` runs).
  - `log.Panic` logs the message and then panics — can be recovered with `recover` if you need to.
  - These calls are commented out in the demo so the rest of the example can run.

### How to run

From the `stdlib/log` directory:

```bash
go run log-demo.go
```

### Docs

For more details, see the official `log` package documentation: `https://pkg.go.dev/log`.

