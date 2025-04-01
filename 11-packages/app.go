package main

import (
	"errors"
	"fmt"
)

func userInput(message string) (value string) {
	fmt.Print(message)
	fmt.Scan(&value)
	return value
}

func userCredentials(message string) (string, string) {
	fmt.Println(message)
	var username string = userInput("Type your email: ")
	var password string = userInput("Type your password: ")
	return username, password
}

func signUp() {
	username, password := userCredentials("Create a new account first!")
	setUsernameToFile(username)
	setPasswordToFile(password)
}

func errorHandling(message string) {
	errorObject := errors.New(message)
	fmt.Println(errorObject)
}

func login(localUsername, localPassword string) {
	username, password := userCredentials("Create a new account first!")

	if username != localUsername {
		errorHandling("The email is incorrect")
		return
	}

	if password != localPassword {
		errorHandling("The password is incorrect")
		return
	}

	fmt.Println("You successfully logged in!")
}

func main() {
	fmt.Println("Welcome to the Go Course Website. Login tu get access to the full course")
	localUsername, username_error := getUsernameFromFile()
	localPassword, _ := getPasswordFromFile()

	if username_error != nil {
		signUp()
	} else {
		login(localUsername, localPassword)
	}
}
