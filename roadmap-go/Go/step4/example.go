package main

import "fmt"

type taxSystem interface {
	calculate() int
}
type RusTax struct {
	income        int
	taxPercentage int
}

func (i RusTax) calculate() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type UsaTax struct {
	income        int
	taxPercentage int
}

func (i UsaTax) calculate() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	RusTax := RusTax{income: 1000, taxPercentage: 10}
	UsaTax := UsaTax{income: 2000, taxPercentage: 20}
	tax := []taxSystem{RusTax, UsaTax}
	result := cal(tax)
	fmt.Println(result)
}
func cal(i []taxSystem) int {
	totalTax := 0
	for _, t := range i {
		totalTax += t.calculate()
	}
	return totalTax
}
