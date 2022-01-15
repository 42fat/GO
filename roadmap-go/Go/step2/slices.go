package main

import "fmt"

func main() {
	s := []int{}
	fmt.Println(len(s), cap(s), s)

	numbers := [5]int{1, 2, 3, 4, 5}
	one := numbers[0:3]
	fmt.Println("Создали слайс с помошью массива", one)

	numbers[1] = 15
	fmt.Println("Слайс изменился вместе с массивом", one)

	two := make([]int, 3, 5)
	fmt.Println("Создали слайс с помошью make", len(two), cap(two), two)

	tree := new([]int)
	fmt.Println("Создали слайс с помошью new", len(*tree), cap(*tree), *tree)

	four := append(two, 15, 16, 16, 74, 32, 231)
	fmt.Println("Слайс ", len(two), cap(two), two)
	fmt.Println("Слайс append", len(four), cap(four), four)

	five := append(two, one...)
	fmt.Println("Слайс из двух слайсов", len(five), cap(five), five)

	six := make([]int, 5)
	k := copy(six, five)
	fmt.Println("Копия слайса ", len(six), cap(six), six, k)

	seven := make([][]int, 2)
	seven[0] = []int{1, 2, 3}
	seven[1] = []int{3, 2, 1}

}
