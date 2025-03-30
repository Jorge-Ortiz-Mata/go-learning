package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 1.3
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	const years = 10.0
	const expectedReturnRate = 5.5

	// This is for the user to input the investment amount
	fmt.Scan(&investmentAmount)

	// Calculate the future value of the investment
	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
