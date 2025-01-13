package main

import "fmt"

func main(){
// 	Your Assignment (Step 1)
// Install Go and set up your environment.
// Write a program to print:
// Your name
// Your favorite programming language
// Your goal for learning Go

	//simple print statemenst
	fmt.Println("My Name is Nahom")
	fmt.Println("My favorite programming lanuage is C++")
	fmt.Println("My goal for learning go is to pass my course")

	//variables
	var name string = "Nahom"
	var age int = 20
	name2 := "Nahom"
	age2 := 20
	fmt.Println(name, age, name2, age2)
	//more var
	var a int = 10
	var b string = "Go"
	var c bool //defualts to false
	//using shorthand
	x := 3.14
	y := "Hello"
	fmt.Println(a, b, c, x, y)

	//If variable is declared without without using, Go initilizes it to zero
	//numeric(int, float) : 0
	//string ""(empty string)
	//bool false

// Write a program to declare:
// Your name using var.
// Your age using shorthand :=.
// A variable for "are you enjoying Go" (as a boolean), but don’t assign a value. Let it use the default value.
// Print the values of all variables.

   var myName string = "Nahom"
   myAge := 20
   var enjoyinGo bool
   fmt.Println("Name: " , myName)
   fmt.Println("Age: " ,myAge)
   if enjoyinGo {
	fmt.Println("I am enjoying Go")
   }else{
	fmt.Println("I am not enjoying Go")
   }
}