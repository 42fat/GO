package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Пополнение баланса", balance, "на сумму", value)
	balance += value
	mutex.Unlock()
	wg.Done()
}
func draw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Снятие с баланса", balance, "на сумму", value)
	balance -= value
	mutex.Unlock()
	wg.Done()
}
func main() {
	balance = 1000

	var wg sync.WaitGroup

	wg.Add(2)

	go deposit(500, &wg)
	go draw(1000, &wg)
	wg.Wait()
	fmt.Println("Balance", balance)
}
