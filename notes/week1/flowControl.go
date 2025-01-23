package main

import "fmt"

func main() {
	// Control Structures in Go

	// 1. Conditional Statements

	// Example 1: if-else
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Example 2: if with declaration
	if num := 10; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Example 3: switch statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	// 2. Loops

	// Example 1: Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// Example 2: Infinite loop with break
	for {
		fmt.Println("This runs forever unless broken")
		break
	}

	// Example 3: Using range
	names := []string{"Nahom", "Alice", "Bob"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	// Example 4: Break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip when i is 2
		}
		if i == 4 {
			break // Exit the loop when i is 4
		}
		fmt.Println(i)
	}

	// Assignment: Control Structures
	// 	Assignment: Control Structures
	// Write a program that:
	// Takes an integer variable age.
	// Prints whether the person is a child, teenager, or adult using if-else.
	// Uses a for loop to print numbers from 1 to age.
	// Uses a switch to categorize the person’s age group:
	// 0–12: "Child"
	// 13–19: "Teenager"
	// 20+: "Adult"

	// 1. Check age group using if-else
	var theAge int = 20
	if theAge >= 0 && theAge <= 12 {
		fmt.Println("You are a Child")
	} else if theAge >= 13 && theAge <= 19 {
		fmt.Println("You are a Teenager")
	} else {
		fmt.Println("You are an Adult")
	}

	// 2. Print numbers from 1 to age
	for i := 1; i <= theAge; i++ {
		fmt.Println("Count:", i)
	}

	// 3. Categorize age group using switch
	switch {
	case theAge >= 0 && theAge <= 12:
		fmt.Println("You are a Child")
	case theAge >= 13 && theAge <= 19:
		fmt.Println("You are a Teenager")
	case theAge > 19:
		fmt.Println("You are an Adult")
	default:
		fmt.Println("Invalid age")
	}
}
