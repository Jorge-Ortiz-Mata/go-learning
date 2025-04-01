package main

import (
	"errors"
	"os"
)

const usernameFile = "user.txt"
const passwordFile = "password.txt"

func setUsernameToFile(username string) {
	os.WriteFile(usernameFile, []byte(username), 0644)
}

func setPasswordToFile(password string) {
	os.WriteFile(passwordFile, []byte(password), 0644)
}

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
