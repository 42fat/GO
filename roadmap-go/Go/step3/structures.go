package main

import (
	"fmt"
)

type Export struct {
	Name string
	age  int
}

type student struct {
	name   string
	age    int
	course int
}
type anonimStudent struct {
	string
	int
}
type class struct {
	student
}

func main() {
	st1 := student{}
	st2 := student{
		age:    21,
		name:   "Maxim",
		course: 3}
	st3 := student{age: 18, name: "Tom", course: 1}
	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)
	st3.name = "Yoda"
	st3.age = 150
	st3.course = 5

	stAdr := &st3
	fmt.Println(stAdr, *stAdr, &stAdr)

	stPtr := new(student)
	fmt.Println(stPtr, *stPtr, &stPtr)

	fmt.Printf("%+v\n", *stAdr)
	fmt.Printf("%v\n", *stAdr)
	fmt.Printf("%#v\n", *stAdr)

	anSt := anonimStudent{}
	anSt.string = "noname"
	anSt.int = 15
	fmt.Println(anSt)

	cl1 := class{st1}
	cl2 := class{}
	cl2.name = cl1.name
	cl2.age = cl1.age
	cl2.course = cl1.course
	fmt.Println(cl2)
	cl2.student.name = cl1.name
	cl2.student.age = cl1.age
	cl2.student.course = cl1.course
	fmt.Println(cl2)

	exp := Export{"Hill", 16}
	exp.Name = "14"
	exp.age = 15

	if st2 == st3 {
		fmt.Println("Они равны")
	} else {
		fmt.Println("Они не равны")
	}
}
