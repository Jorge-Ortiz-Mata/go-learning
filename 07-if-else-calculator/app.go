package main

import "fmt"

func main() {
	var result float64 = 0
	var firstValue float64
	var secondValue float64
	var option int

	fmt.Print("Insert the operation you want to do: ")
	fmt.Scan(&option)

	fmt.Print("Insert the first value: ")
	fmt.Scan(&firstValue)
	fmt.Print("Insert the second value: ")
	fmt.Scan(&secondValue)

	if option == 1 {
		result = sum(firstValue, secondValue)
	} else if option == 2 {
		result = substract(firstValue, secondValue)
	} else if option == 3 {
		result = multiply(firstValue, secondValue)
	} else {
		result = divide(firstValue, secondValue)
	}

	fmt.Printf("The result of the operation is: %.2f\n", result)
}

func sum(firstValue float64, secondValue float64) float64 {
	return firstValue + secondValue
}

func substract(firstValue float64, secondValue float64) float64 {
	return firstValue - secondValue
}

func multiply(firstValue float64, secondValue float64) float64 {
	return firstValue * secondValue
}

func divide(firstValue float64, secondValue float64) float64 {
	return firstValue / secondValue
}
