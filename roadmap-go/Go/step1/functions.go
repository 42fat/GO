package main

import "fmt"

func one(num int) string {
	return "Я не могу быть вызвана в другом файле"
}

func Two(num int) string {
	return "Я  могу быть вызвана в другом файле"
}
func Tree() {
	fmt.Println("Я ничего не возврашаю и ничего не принимаю")
	return
}
func Four(num int) (int, int) {
	fmt.Println("Я возвращаю два int")
	return num, num
}
func five(num int) (res int) {
	res = num * 2
	return res
}
func six(num *int) {
	*num = *num * 6
	return
}

func nine(num int) func(int) string {
	return Two
}

func main() {
	seven := func(a, b int) int {
		return a % b
	}

	func(name string) {
		fmt.Printf("Немедленная, %s", name)
	}("функция")

	fmt.Println(one(1))

	fmt.Println(Two(1))

	Tree()

	i, _ := Four(1)
	m, j := Four(1)
	fmt.Println(i, j, m)

	k := five(i)
	fmt.Println(k)

	six(&k)
	fmt.Println(k)

	res := seven(k, m)
	fmt.Println(res)

	f := nine(1)
	fmt.Println(f(1))
}
