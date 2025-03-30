package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	// Calculate the future value of the investment
	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

	fmt.Println(futureValue)
}
