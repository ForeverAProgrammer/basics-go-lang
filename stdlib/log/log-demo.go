package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// --- All log flag constants (combine with | to mix) ---
	//
	//   log.Ldate         2009/11/10
	//   log.Ltime                    23:00:00
	//   log.Lmicroseconds            23:00:00.000000  (overrides Ltime)
	//   log.Llongfile     /full/path/to/file.go:42
	//   log.Lshortfile    file.go:42               (overrides Llongfile)
	//   log.LUTC          use UTC instead of local time zone
	//   log.Lmsgprefix    move prefix to just before message (not start of line)
	//   log.LstdFlags     Ldate | Ltime  (the default)

	// By default, the standard logger writes to stderr with a date+time prefix.
	log.Println("Default logger: this is an info-style message")
	log.Printf("Formatted number: %d", 42)

	// Customize prefix and flags on the default logger.
	log.SetPrefix("INFO  ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With prefix and short file flag")

	// --- log.Lmsgprefix ---
	// By default the prefix appears at the very start of the line, before the
	// date/time. Lmsgprefix moves it to just before the message text.
	//
	//   Without Lmsgprefix:  INFO  2009/11/10 23:00:00 message
	//   With    Lmsgprefix:  2009/11/10 23:00:00 INFO  message
	msgPrefixLogger := log.New(os.Stdout, "INFO  ", log.LstdFlags|log.Lmsgprefix)
	msgPrefixLogger.Println("Lmsgprefix: prefix appears just before the message")

	// Create separate loggers with different "types" via prefixes.
	infoLogger := log.New(os.Stdout, "INFO  ", log.LstdFlags)
	warnLogger := log.New(os.Stdout, "WARN  ", log.LstdFlags)
	errorLogger := log.New(os.Stderr, "ERROR ", log.LstdFlags)

	infoLogger.Println("Application started")
	warnLogger.Println("This is a warning example")
	errorLogger.Println("This is an error example")

	// --- Write to a file ---
	// log.New accepts any io.Writer, so you can log directly to a file.
	f, err := os.OpenFile("demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer f.Close()

	fileLogger := log.New(f, "FILE  ", log.LstdFlags)
	fileLogger.Println("This line is written to demo.log")
	fileLogger.Printf("Formatted: value=%d", 99)

	// --- io.MultiWriter: log to file and stderr simultaneously ---
	// io.MultiWriter fans writes out to multiple io.Writers at once.
	// This is the standard way to mirror log output to both a file and the console.
	multi := io.MultiWriter(f, os.Stderr)
	multiLogger := log.New(multi, "MULTI ", log.LstdFlags)
	multiLogger.Println("This line goes to both demo.log and stderr")

	// --- log.SetFlags(0): suppress all prefix fields ---
	// Passing 0 removes the timestamp and file info entirely, leaving just the
	// message. Useful for CLIs or tests where timestamps are noise.
	log.SetFlags(0)
	log.SetPrefix("")
	log.Println("No timestamp or file info")

	// log.Fatal and log.Panic also exist:
	//
	//   log.Fatal("fatal error")   // logs the message then os.Exit(1)
	//   log.Panic("panic error")   // logs the message then panics
	//
	// They stop the program, so they are commented out here to let the
	// rest of the demo run. In real programs, use them sparingly when the
	// program truly cannot continue.
}
