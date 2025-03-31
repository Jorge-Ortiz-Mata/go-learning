package main

import (
	"errors"
	"fmt"
	"os"
)

const usernameFile = "user.txt"
const passwordFile = "password.txt"

func getUsernameFromFile() (string, error) {
	data, err := os.ReadFile(usernameFile)
	if err != nil {
		return "", errors.New("There is not username created")
	}

	username := string(data)
	return username, nil
}

func getPasswordFromFile() (string, error) {
	data, err := os.ReadFile(passwordFile)
	if err != nil {
		return "", errors.New("There is not password created")
	}

	password := string(data)
	return password, nil
}

func setUsernameToFile(username string) {
	os.WriteFile(usernameFile, []byte(username), 0644)
}

func setPasswordToFile(password string) {
	os.WriteFile(passwordFile, []byte(password), 0644)
}

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
