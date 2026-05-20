## Log package demo

This example shows how to use Go's built-in `log` package for basic logging.

### What it covers

- **Default logger**
  - Using `log.Println` and `log.Printf`.
  - Understanding that the default logger writes to **stderr** with date/time by default.
- **Custom prefixes and flags**
  - Changing the prefix with `log.SetPrefix`.
  - Changing output format with `log.SetFlags` (e.g. `log.LstdFlags`, `log.Lshortfile`).
- **All flag constants**
  - `log.Ldate`, `log.Ltime`, `log.Lmicroseconds`, `log.Llongfile`, `log.Lshortfile`, `log.LUTC`, `log.Lmsgprefix`, `log.LstdFlags`.
  - Flags can be combined with `|`.
- **`log.Lmsgprefix`**
  - By default the prefix appears at the very start of the line, before the date/time.
  - `Lmsgprefix` moves it to just before the message text.
- **Different "log types" via prefixes**
  - Creating multiple `*log.Logger` instances with `log.New`, such as:
    - `INFO` logger writing to stdout.
    - `WARN` logger writing to stdout.
    - `ERROR` logger writing to stderr.
- **Writing to a file**
  - `log.New` accepts any `io.Writer`, so you can pass an `*os.File` to log directly to disk.
  - Use `os.OpenFile` with `os.O_CREATE|os.O_WRONLY|os.O_APPEND` to open or create a log file.
- **`io.MultiWriter`: log to multiple destinations at once**
  - `io.MultiWriter` fans writes out to multiple `io.Writer` values simultaneously.
  - Standard pattern for mirroring output to both a file and the console.
- **`log.Fatal` and `log.Panic`**
  - `log.Fatal` logs the message and then calls `os.Exit(1)` — the program stops immediately (no `defer` runs).
  - `log.Panic` logs the message and then panics — can be recovered with `recover` if you need to.
  - These calls are commented out in the demo so the rest of the example can run.

### How to run

From the `stdlib/log` directory:

```bash
go run log-demo.go
```

A file named `demo.log` will be created in the same directory showing the output
from the file logger and the multi-writer logger.

### Docs

For more details, see the official `log` package documentation: `https://pkg.go.dev/log`.
