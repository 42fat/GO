package main

import (
	"fmt"
	"os"
)

func main() {
	file, error := os.Open("file.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(file.Name(), "open")
	}
}
