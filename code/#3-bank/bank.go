package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceString), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000.0, errors.New("file does not exist  ")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse stored balance")
	}

	return balance, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")   
		// panic("can't continue sorry it's is important error") // crashes the entire application
	}

	fmt.Println("-----------------------------")
	fmt.Println("Welcome to go Bank!")
	fmt.Println("-----------------------------")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("-----------------------------")

		fmt.Println("1. Check The Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)

		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new account balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance. Your account balance is:", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new account balance is:", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using go Bank!")
			return
		}
	}
}
