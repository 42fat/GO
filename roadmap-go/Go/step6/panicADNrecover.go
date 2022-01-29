package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	a := []string{"a", "b"}
	myprint(a, 0)
	panic("break")
	fmt.Println("Normal exit")
}

func myprint(arr []string, index int) {
	if index < 0 || index > len(arr) {
		panic("Out of bound")
	}
	fmt.Println(arr[index])
}

func print(arr []string, index int) {
	fmt.Println(arr[index])
}
