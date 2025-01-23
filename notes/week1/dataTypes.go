package main

import "fmt"

func main() {
	// Constants: Values that cannot be changed
	const Pi = 3.14
	const Greeting = "Hello, Go"

	// Grouped constants
	const (
		Pi2       = 3.14
		Greeting2 = "Hello, World"
		MaxAge    = 120
	)

	// Typed and Untyped Constants
	const TypedPi float64 = 3.14 // Typed constant
	const UntypedPi = 3.12       // Untyped constant

	// Assignment: Constants and Data Types
	// Declare constants for:
	// - Your full name
	// - The value of Pi (3.14159)
	// - The maximum number of students in a classroom
	const (
		name        string  = "Nahom"
		Pii         float32 = 3.14159
		maxStudents         = 25
	)

	// Declare variables for:
	// - Your age (integer)
	// - Your current grade in Go (as a float between 0 and 100)
	// - A boolean indicating if you like constants
	const (
		MyAge         int16 = 20
		CurrentGrade  int8  = 10
		likeConstants bool  = true
	)

	// Print all values
	fmt.Println("Name:", name)
	fmt.Println("Pi:", Pii)
	fmt.Println("Max Students:", maxStudents)
	fmt.Println("Age:", MyAge)
	fmt.Println("Current Grade:", CurrentGrade)
	fmt.Println("Likes Constants:", likeConstants)
}

// Notes on Integer Types and Their Ranges:
// int8   : 1 byte  -> -128 to 127
// uint8  : 1 byte  -> 0 to 255
// int16  : 2 bytes -> -32,768 to 32,767
// uint16 : 2 bytes -> 0 to 65,535
// int32  : 4 bytes -> -2,147,483,648 to 2,147,483,647
// uint32 : 4 bytes -> 0 to 4,294,967,295
// int64  : 8 bytes -> -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// uint64 : 8 bytes -> 0 to 18,446,744,073,709,551,615
