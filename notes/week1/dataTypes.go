package main

import "fmt"


func main(){
	//Const value that cannot be changed 
	const Pi = 3.14
	const Greeting = "hello, Go"

	const(
		Pi2 = 3.14
		Greeting2 = "Hello world"
		MaxAge = 120
	)

	const TyepedPi float64 = 3.14 //typed
	const unTypedPi = 3.12 //untyped

// Assignment: Constants and Data Types
// Declare constants for:
// Your full name.
// The value of Pi (3.14159).
// The maximum number of students in a classroom.
// Declare variables for:
// Your age (integer).
// Your current grade in Go (as a float between 0 and 100).
// A boolean indicating if you like constants.
// Print all values.

const(
	name string = "Nahom"
	Pii float32 = 3.14159
	maxStudents = 25
)

const(
	MyAge int16 = 20
	CurrentGrade int8 = 10
	likeConstants bool  = true
)

    fmt.Println("Name:", name)
    fmt.Println("Pi:", Pii)
    fmt.Println("Max Students:", maxStudents)
    fmt.Println("Age:", MyAge)
    fmt.Println("Current Grade:", CurrentGrade)
    fmt.Println("Likes Constants:", likeConstants)
}