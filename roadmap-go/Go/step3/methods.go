package main

import (
	"fmt"
)

type test struct {
	name string
	age  int
}

func (st test) info() {
	fmt.Println(st)
}
func (st test) getName() string {
	return st.name
}
func (st *test) setName(name string) {
	st.name = name
}
func main() {
	tmp := test{name: "Max", age: 16}
	tmp.info()
	fmt.Println(tmp.getName())
	tmp.setName("Tom")
	fmt.Println(tmp)
}
