package main

import "fmt"

var num = 1
var tnum int = 1

//variable := 1 , мы не можем использовать
func main() {
	var variableType string = "var variableType"
	var variable = "var variable"
	superVariable := 1
	num = 6
	tnum = 5
	fmt.Println(variable, variableType, superVariable, num, tnum)
}
