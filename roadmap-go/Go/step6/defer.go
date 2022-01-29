package main

import "fmt"

func Hello() {
	fmt.Println("Hello")
}
func world() {
	fmt.Println("world")
}
func end() {
	fmt.Println("!")
}

func main() {
	defer Hello()
	defer world()
	end()

	a := 15

	defer fmt.Println("Значение запомнится ", a)

	a = 44

}
