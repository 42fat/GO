package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go one(a)
	go two(b)

	select {
	case msg1 := <-a:
		fmt.Println(msg1)
	case msg2 := <-b:
		fmt.Println(msg2)
	case b <- "Hello boys":
	case <-time.After(time.Second * 1):
	default:
		fmt.Println("Error")
	}
	ch := make(chan int, 10)
	arr(ch)
	printarr(ch)
}
func printarr(info chan int) {
	for {
		select {
		case msg1 := <-info:
			fmt.Println("id:", msg1)
		case <-time.After(time.Second * 2):
			return
		}
	}
}

func arr(info chan<- int) {
	i := 0
	for i < 10 {
		info <- i
		i++
	}
}
func one(str chan string) {
	str <- "one"
}

func two(str chan string) {
	str <- "two"
}
