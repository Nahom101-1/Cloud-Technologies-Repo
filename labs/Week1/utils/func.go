package utils

// Create function type with given signature
type MyFunction func(a, b interface{}) (interface{}, error)

/*
Variable holding the function
*/
var internalFunction func(string, int) string

/*
Corresponding setter and getters
*/
func SetInternalFunction(f func(string, int) string) {
	internalFunction = f
}

func GetInternalFunction() func(string, int) string {
	return internalFunction
}

/*
Interface{} variant of function
*/
var anotherInternalFunction MyFunction

func SetAnotherFunction(f MyFunction) {
	anotherInternalFunction = f
}

func GetAnotherFunction() MyFunction {
	return anotherInternalFunction
}
