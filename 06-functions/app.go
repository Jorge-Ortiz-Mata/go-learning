package main

import (
	"fmt"
)

func main() {
	var nameReceived = setName()
	fmt.Println(nameReceived)
	const myAge int = 20
	sayYourAge(myAge)
}

func setName() string {
	return "Hello, I am Jorge Ortiz."
}

func sayYourAge(age int) {
	fmt.Printf("Hello. I am %v\n", age)
}
