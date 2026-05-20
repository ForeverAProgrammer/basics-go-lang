package main

import (
	"log"
	"os"
)

func main() {
	// By default, the standard logger writes to stderr with a date+time prefix.
	log.Println("Default logger: this is an info-style message")
	log.Printf("Formatted number: %d", 42)

	// Customize prefix and flags on the default logger.
	log.SetPrefix("INFO  ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With prefix and short file flag")

	// Create separate loggers with different "types" via prefixes.
	infoLogger := log.New(os.Stdout, "INFO  ", log.LstdFlags)
	warnLogger := log.New(os.Stdout, "WARN  ", log.LstdFlags)
	errorLogger := log.New(os.Stderr, "ERROR ", log.LstdFlags)

	infoLogger.Println("Application started")
	warnLogger.Println("This is a warning example")
	errorLogger.Println("This is an error example")

	// log.Fatal and log.Panic also exist:
	//
	//   log.Fatal("fatal error")   // logs the message then os.Exit(1)
	//   log.Panic("panic error")   // logs the message then panics
	//
	// They stop the program, so they are commented out here to let the
	// rest of the demo run. In real programs, use them sparingly when the
	// program truly cannot continue.
}

