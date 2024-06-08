package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax float64

	revenue, expenses, tax = IO("Revenue"), IO("Expenses"), IO("Tax")

	EBT, EAT := calc(revenue, expenses, tax)

	fmt.Println("EBT: ", EBT)
	fmt.Println("EAT: ", EAT)
}

func IO(text string) (value float64) {
	fmt.Print(text, ": ")
	fmt.Scan(&value)
	return value
}

func calc(revenue, expenses, tax float64) (EBT, EAT float64) {
	EBT = revenue - expenses
	EAT = revenue*(1-tax) - expenses
	return EBT, EAT
}
