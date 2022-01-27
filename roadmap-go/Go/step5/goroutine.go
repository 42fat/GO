package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go start()
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		go id(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumCPU())
}

func start() {
	go end()
	fmt.Println("Hello world")
}
func end() {
	fmt.Println("Goodbye")
}

func id(num int) {
	fmt.Println("id:", num)
}
