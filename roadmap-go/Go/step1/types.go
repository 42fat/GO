package main

import "fmt"

// to run noname.go run go noname.go
func main() {
	var integer int
	var floatSmall float32
	var floatBig float64
	var boolean bool
	var str string
	var complex complex64
	var let rune
	fmt.Println("Значение по умолчанию всех переменных")
	fmt.Printf("%T ", integer)
	fmt.Println("\"", integer, "\"")

	fmt.Printf("%T ", floatSmall)

	fmt.Println("\"", floatSmall, "\"")
	fmt.Printf("%T ", floatBig)
	fmt.Println("\"", floatBig, "\"")

	fmt.Printf("%T ", boolean)
	fmt.Println("\"", boolean, "\"")

	fmt.Printf("%T ", str)
	fmt.Println("\"", str, "\"")

	fmt.Printf("%T ", complex)
	fmt.Println("\"", complex, "\"")

	fmt.Printf("%T ", let)
	fmt.Println("\"", let, "\"")
}
