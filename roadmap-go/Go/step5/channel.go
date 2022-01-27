package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go send(a)
	go print(a)
	time.Sleep(2 * time.Second)

	b := make(chan int, 10)
	go sendmass(b)
	go printmass(b)
	time.Sleep(2 * time.Second)

	fmt.Println("Размер а :", cap(a), "Размер б", cap(b))
	fmt.Println("Длина", len(b))

	ch := make(chan string, 1)
	ch <- "end"
	val, ok := <-ch
	fmt.Printf("Val: %d OK: %t\n", val, ok)

	close(ch)
	val, ok = <-ch
	fmt.Printf("Val: %d OK: %t\n", val, ok)
}

func print(info chan int) {
	val := <-info
	fmt.Println("print", val)
}
func send(info chan int) {
	fmt.Println("Send")
	a := 123
	info <- a
}
func sendmass(info chan<- int) {
	fmt.Println("Send")
	i := 0
	for i <= 10 {
		info <- i
		i++
	}
}

func printmass(info <-chan int) {
	fmt.Println("Print")
	i := 0
	for i <= 10 {
		fmt.Println(<-info)
		i++
	}
}
