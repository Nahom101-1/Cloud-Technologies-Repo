package main

import (
	"labs/Week1/utils"
	"strconv"
)


func main(){
	utils.SetInternalFunction(func(a string, b int) string {
	return a + " has the age of " + strconv.Itoa(b)
})

}
