package main

import "fmt"

const a = 125

func main() {
	const str string = "Hello world const"
	const pi = 3.14
	const (
		e     = 2.7
		ename = "exp"
	)
	const a = 666
	fmt.Println(a)
	fmt.Println(str)
}
