package utils

import "fmt"

// Create function type with given signature
type MyFunction func(a, b interface{}) (interface{}, error)

//	var defaultFunction(name string, age int) string{
//		return "Function not"
//	}
//
// Variable holding the function
var internalFunction func(string, int) string

/*
Corresponding stter and getters
*/
func SetInternalFunction(f func(string, int) string) {
	if f == nil {
		fmt.Println("Error: cannot assign a nul function!")
		return
	}
	internalFunction = f
}

func GetInternalFunction() func(string, int) string {
	return internalFunction
}

// var anotherInternalFunction MyFunction

// /*
// Interface{} variant of function
// */
// func SetAnotherFunction(f MyFunction){
// 	anotherInternalFunction = f
// }

// func getAnotherFunction() MyFunction {
// 	return anotherInternalFunction
// }
