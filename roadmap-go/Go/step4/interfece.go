package main

import "fmt"

type animal interface {
	breathe()
	walk()
	age() int
}
type lion struct {
	age_ int
}

func (l lion) age() int {
	return l.age_
}
func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age_ int
}

func (l dog) breathe() {
	fmt.Println("Dog breathes")
}

func (l dog) walk() {
	fmt.Println("Dog walk")
}

func (l dog) age() int {
	return l.age_
}

func main() {
	var a animal
	a = lion{10}
	a.walk()
	a.breathe()
	fmt.Println(a.age())
	fmt.Println(CallAge(a))
	var b animal
	b = dog{15}
	b.walk()
	b.breathe()
	fmt.Println(b.age())
	fmt.Println(CallAge(b))
}

func CallAge(a animal) int {
	return a.age()
}
