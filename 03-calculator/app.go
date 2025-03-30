package main

import (
	"fmt"
)

func main() {
	var firstValue float64
	var secondValue float64

	fmt.Print("Type the first value: ")
	fmt.Scan(&firstValue)
	fmt.Print("Type the second value: ")
	fmt.Scan(&secondValue)

	var result float64 = firstValue + secondValue

	// %.2f is a format specifier that formats the number to 2 decimal places
	fmt.Printf("The operation result is: %.2f\n", result)
}
