package main

import "fmt"

func main() { 
	accountBalance  := 1000.0

	fmt.Println("-----------------------------")
	fmt.Println("Welcome to go Bank!")
	fmt.Println("-----------------------------")
	fmt.Println("What would you like to do?")
	fmt.Println("-----------------------------")

	// options
	fmt.Println("1. Check The Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choise int
	fmt.Print("Enter your choise: ")
	fmt.Scan(&choise)

	// wantsCheckBalance := choise == 1
	// wantsDepositMoney := choise == 2
	// wantsWithdrawMoney := choise == 3
	// wantsExit := choise == 4

	if choise == 1 {
		fmt.Println("Your account balance is:", accountBalance)
	} else if choise == 2 {
		fmt.Print("Your deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. must be greater than 0");
			return;
		}
 
		accountBalance += depositAmount
		fmt.Println("Your new account balance is:", accountBalance)
	} else if choise == 3 {
		fmt.Print("Your withdraw amount:  ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. must be greater than 0");
			return;
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient balance. Your account balance is:", accountBalance)
			return;
		}

		accountBalance -= withdrawAmount
		fmt.Println("Your new account balance is:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}   