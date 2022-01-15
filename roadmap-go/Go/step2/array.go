package main

import "fmt"

func main() {
	const sizeOfArray = 4
	sample := [sizeOfArray]int{1, 2, 6, 3}
	fmt.Println("Элементы массива ", sample[0], sample[1], sample[2], sample[3])

	one := [...]int{1, 2, 3}
	fmt.Println("Массив без указания размера", one[0], one[1], one[2])

	two := [3]int{}
	fmt.Println("Массив без указания элементов", two[0], two[1], two[2])

	fmt.Println("Массив можно вывести так", one, "и длина его равна", len(one))

	two = one
	fmt.Println("Массивы ", one, two)
	two[2] = 15
	fmt.Println("Массивы не связаны ", one, two)

	tree := [2][3]int{}
	tree[0] = one
	tree[1] = two
	fmt.Println("Двухмерные массив", tree, "i:", len(tree[0]), "j:", len(tree))
}
