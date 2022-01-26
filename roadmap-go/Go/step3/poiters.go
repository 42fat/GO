package main

import "fmt"

func main() {
	a := new(int)
	*a = 10
	fmt.Println("Значение переменной ", *a)
	fmt.Println("Адресс переменной ", a)

	one := 2
	two := &one
	fmt.Println("One ", one)
	fmt.Println("Two ", *two)
	fmt.Println("Adress One ", &one)
	fmt.Println("Adress Two ", two)
	*two = 10
	fmt.Println("One ", one)
	fmt.Println("Two ", *two)

	one = 15
	two = &one
	three := &two
	fmt.Println("One ", one, "Adress one", &one)
	fmt.Println("Two ", *two, "Address one", two)
	fmt.Println("Three ", **three, "Adress one ", *three, "Adress three", three)

}
