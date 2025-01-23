package main

import (
	"errorHandling/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	// Call a utility function that may return an error
	returnA, err := utils.FunctionThatThrowsErrors("Some input", "some other input")
	if err != nil {
		// Write error to standard error manually
		// os.Stderr.WriteString writes the error message directly to the error stream.
		// It doesn't add timestamps or automatic newlines.
		// Useful for CLI tools where error messages need to be explicitly separated.
		os.Stderr.WriteString(err.Error() + "\n")

		// Use the log package to print errors with timestamps
		// log.Println automatically appends a newline and logs to standard error with a timestamp.
		log.Println(err)

		// Use log.Printf to format the error message
		// log.Printf is useful for structured logging, allowing custom messages with context.
		log.Printf("Error occurred: %s - Status Code: %d", err, 500)

		// Log the error and exit the program immediately
		// log.Fatalln logs the error message and exits with a status code of 1.
		// It's commonly used for unrecoverable errors.
		log.Fatalln("Printing error and exiting (with exit code 1):", err.Error())

		// Log the error and panic
		// log.Panicln logs the error message and then panics, which is useful for debugging.
		// Avoid using panics in production as they abruptly terminate the program unless recovered.
		log.Panicln("Panicking out of code based on error:", err.Error())

		// This line is unreachable because log.Fatalln and log.Panicln terminate the program.
		fmt.Println("This will never print")
	}

	// This line will only execute if no error occurred.
	fmt.Println("Success: " + returnA)
}
