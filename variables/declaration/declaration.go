package main

import (
	"fmt"
	"reflect"
)

func main() {
	// --- Naming rules ---
	// Exported — starts with a capital letter, accessible from other packages:
	//   var MaxRetries int = 3
	//   var HTTPClient  = ...
	//
	// Unexported — starts with a lowercase letter, only accessible within this package:
	//   var maxRetries int = 3
	//   var httpClient  = ...
	//
	// Community conventions (https://go.dev/doc/effective_go#names):
	//   - camelCase, not snake_case:         firstName      not  first_name
	//   - Acronyms stay uppercase:            userID, HTTPServer  not  userId, HttpServer
	//   - Short names for short-lived vars:   i, n, err, ok
	//   - Boolean names read naturally:       isActive, hasError, canDelete
	//   - Avoid stutter with package name:    user.Name      not  user.UserName

	// --- Declaring a variable twice (compile-time error) ---
	// Each variable can only be declared once within the same scope.
	// Uncommenting any line below will prevent the program from building:
	//
	//   var age int = 99     // age redeclared in this block
	//   score := 200         // no new variables on left side of :=
	//
	// Note: := requires at least one new variable on the left side.
	// It can be used again only if it introduces a new variable alongside an existing one:
	//
	//   score, newVar := 200, true   // valid — newVar is new

	// --- Zero Values ---
	// Variables declared without a value are automatically assigned their zero value
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println("Zero Values:")
	fmt.Printf("int: %d | type: %s\n", i, reflect.TypeOf(i))
	fmt.Printf("float64: %f | type: %s\n", f, reflect.TypeOf(f))
	fmt.Printf("string: '%s' | type: %s\n", s, reflect.TypeOf(s))
	fmt.Printf("bool: %t | type: %s\n", b, reflect.TypeOf(b))
	fmt.Println()

	// --- Single-line var (explicit type) ---
	var age int = 30
	var username string = "Alice"

	fmt.Println("Single-line var (explicit type):")
	fmt.Printf("age: %d | type: %s\n", age, reflect.TypeOf(age))
	fmt.Printf("username: '%s' | type: %s\n", username, reflect.TypeOf(username))
	fmt.Println()

	// --- Single-line var (inferred type) ---
	// If you assign a value at the same time as you declare a variable, you can
	// omit the type — Go will use the type of the assigned value instead:
	//
	//   var score int = 100  →  var score = 100   (type inferred as int)
	//   var label string = "beginner"  →  var label = "beginner"  (type inferred as string)
	var score = 100
	var label = "beginner"

	fmt.Println("Single-line var (inferred type):")
	fmt.Printf("score: %d | type: %s\n", score, reflect.TypeOf(score))
	fmt.Printf("label: '%s' | type: %s\n", label, reflect.TypeOf(label))
	fmt.Println()

	// --- Grouped var block ---
	// Idiomatic for declaring several related variables together
	var (
		city    string = "London"
		country string = "UK"
		pop     int    = 9000000
	)

	fmt.Println("Grouped var block:")
	fmt.Printf("city: '%s' | type: %s\n", city, reflect.TypeOf(city))
	fmt.Printf("country: '%s' | type: %s\n", country, reflect.TypeOf(country))
	fmt.Printf("pop: %d | type: %s\n", pop, reflect.TypeOf(pop))
	fmt.Println()

	// --- Short assignment := ---
	// Most common form inside functions — type is always inferred
	// Note: := can only be used inside a function, not at package level
	x := 42
	y := 3.14
	name := "Go"
	isActive := true

	fmt.Println("Short assignment :=")
	fmt.Printf("x: %d | type: %s\n", x, reflect.TypeOf(x))
	fmt.Printf("y: %f | type: %s\n", y, reflect.TypeOf(y))
	fmt.Printf("name: '%s' | type: %s\n", name, reflect.TypeOf(name))
	fmt.Printf("isActive: %t | type: %s\n", isActive, reflect.TypeOf(isActive))
	fmt.Println()

	// --- Multi-variable short assignment ---
	lang, year, active := "Go", 2024, false

	fmt.Println("Multi-variable short assignment:")
	fmt.Printf("lang: '%s' | type: %s\n", lang, reflect.TypeOf(lang))
	fmt.Printf("year: %d | type: %s\n", year, reflect.TypeOf(year))
	fmt.Printf("active: %t | type: %s\n", active, reflect.TypeOf(active))
	fmt.Println()

	// --- Multi-variable assignment ---
	// Assigning new values to already-declared variables in one line
	// This same syntax is used when a function returns multiple values:
	//   min, max := someFunction()
	length, width := 1.2, 2.4
	fmt.Println("Multi-variable assignment (before):")
	fmt.Printf("length: %f, width: %f\n", length, width)

	length, width = 5.0, 8.5
	fmt.Println("Multi-variable assignment (after):")
	fmt.Printf("length: %f, width: %f\n", length, width)
	fmt.Println()

	// --- Wrong type assignment (compile-time error) ---
	// Go is statically typed — mismatched types are caught at compile time,
	// before the program ever runs. There is no runtime exception to catch.
	//
	// Uncommenting any line below will prevent the program from building:
	//
	//   var n int = "42"   // cannot use "42" (untyped string constant) as int value
	//   var n int = 3.14   // cannot use 3.14 (untyped float constant) as int value
	//   x = "hello"        // cannot use "hello" (untyped string constant) as int value in assignment
}
