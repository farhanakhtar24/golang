package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var inversmentAmount float64 = 1000 // type inference long hand
	// var years float64 = 10 // type inference long hand
	// var expectedReturnRate = 5.5
	// var inversmentAmount, years float64 = 1000, 10

	// expectedReturnRate := 5.5 // type inference short hand
	// inversmentAmount, years := 1000.0, 10.0 // type inference short hand

	inversmentAmount, expectedReturnRate := 1000.0, 5.5
	var years float64

	fmt.Print("Enter inversment amount: ")
	fmt.Scan(&inversmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(inversmentAmount, expectedReturnRate, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value:", futureRealValue)

	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Real Value: %.2f\n", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value:%.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value:%.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func calculateFutureValues(inversmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := inversmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	// return
	return fv, frv
}
