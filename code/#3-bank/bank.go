package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0

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

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using go Bank!")
			return
		}
	}
} 
