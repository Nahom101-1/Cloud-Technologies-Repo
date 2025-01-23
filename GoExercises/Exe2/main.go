package main

import (
	"GoExercises/Exe2/utils"
	"fmt"
	"strings"
)

func Greet(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
}

func Shout(name string, age int) string {
	return fmt.Sprintf("HELLO, %s! YOU ARE %d.", strings.ToUpper(name), age)
}
func main() {
	fn := utils.GetInternalFunction()
	fmt.Println(fn("Alice", 30))
	utils.SetInternalFunction(Greet)
	fn = utils.GetInternalFunction()
	fmt.Println(fn("Alice", 30))
	utils.SetInternalFunction(Shout)
	fn = utils.GetInternalFunction()
	fmt.Println(fn("Alice", 30))
	utils.SetInternalFunction(Greet)
	fn = utils.GetInternalFunction()
	fmt.Println(fn("Bob", 30))
}
