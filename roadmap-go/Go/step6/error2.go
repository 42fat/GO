package main

import (
	"errors"
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error one"
}
func main() {
	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3:%w", e2)
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(errors.Unwrap(e3))

	e4 := e3
	if e4 == e3 {
		fmt.Println("Одинаковые")
	}
	if errors.Is(e2, e1) {
		fmt.Println("Одинаковые")
	}
}
