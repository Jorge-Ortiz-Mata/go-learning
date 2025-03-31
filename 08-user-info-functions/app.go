package main

import "fmt"

func main() {
	var username = getUserInput("username")
	var age = getUserInput("age")
	var email = getUserInput("email")
	fmt.Printf("Your username is: %v\n", username)
	fmt.Printf("Your age is: %v\n", age)
	fmt.Printf("Your email is: %v\n", email)
}

func getUserInput(value string) (inputValue string) {
	fmt.Printf("Type your %v: ", value)
	fmt.Scan(&inputValue)

	return inputValue
}
