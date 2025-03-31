package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFilename string = "database.txt"

func saveBalanceToFile(balance float64) {
	var stringBalance = fmt.Sprint(balance)
	os.WriteFile(balanceFilename, []byte(stringBalance), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFilename)

	if err != nil {
		return 1000.0, errors.New("Failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("Failed to parse data")
	}

	return balance, nil
}

func main() {
	var option int
	balance, err := getBalanceFromFile()

	if err != nil {
		fmt.Print("Something went wrong:")
		return // Remove this line to execute the rest of the code.
	}

	fmt.Println("Welcome to the Bank of America. What would you like to do today?")

	for {
		option = userOptionInput()

		switch option {
		case 1:
			showBalance(balance)
		case 2:
			balance = depositMoney(balance)
			saveBalanceToFile(balance)
		case 3:
			balance = withdrawMoney(balance)
			saveBalanceToFile(balance)
		case 4:
			goodbyeMessage()
		default:
			invalidOption()
		}

		if option == 4 {
			break
		}
	}
}

func showBalance(balance float64) {
	fmt.Printf("You current balance is: %v\n\n", balance)
}

func getUserInput(message string) (value float64) {
	fmt.Println(message)
	fmt.Scan(&value)
	return value
}

func depositMoney(balance float64) float64 {
	var moneyToDeposit float64 = getUserInput("How much money you want to deposit?")

	if moneyToDeposit > 0 {
		var newBalance float64 = balance + moneyToDeposit
		fmt.Printf("You new balance is: %v\n\n", newBalance)
		return newBalance
	} else {
		fmt.Printf("You have to deposit at least 1 dollar\n")
		return balance
	}
}

func withdrawMoney(balance float64) float64 {
	var moneyToWithdraw float64 = getUserInput("How much money you want to withdraw?")

	if moneyToWithdraw > balance {
		fmt.Printf("You have only %v in your bank account.\n\n", balance)
		return balance
	} else {
		var newBalance float64 = balance - moneyToWithdraw
		fmt.Printf("You have withdraw: %v. You current balance is: %v\n\n", moneyToWithdraw, newBalance)
		return newBalance
	}
}

func userOptionInput() (option int) {
	fmt.Println("1. Check my balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Scan(&option)

	return option
}

func goodbyeMessage() {
	fmt.Println("Thanks for choosing us. Have a nice day")
}

func invalidOption() {
	fmt.Println("Invalid Option. Please try again.")
}
