package main

import (
	"fmt"
)

func main() {
	var nameReceived = setName()
	fmt.Println(nameReceived)
	const myAge int = 20
	sayYourAge(myAge)
	var randomNumber1, randomNumber2 = randomNumbers()
	fmt.Println("Random numbers are: ", randomNumber1, randomNumber2)
}

func setName() string {
	return "Hello, I am Jorge Ortiz."
}

func sayYourAge(age int) {
	fmt.Printf("Hello. I am %v\n", age)
}

// In order to return multiple numbers, we should specify the type of each number
func randomNumbers() (int, float64) {
	return 40, 2.20
}
