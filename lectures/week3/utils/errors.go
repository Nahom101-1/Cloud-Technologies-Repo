package utils

import "errors"

func FunctionThatThrowsErrors(a string, b string) (string, error) {
	return "some random string", errors.New("some error message")
}
