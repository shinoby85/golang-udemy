package main

import "fmt"

func main() {
	accountBalance := 10000.0
	fmt.Println("Welcome to Go Bank")

	for {
		var choice int = showMenu()

		if choice == 1 {
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}
			accountBalance += deposit
			fmt.Println("Your balance updated! New amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Your balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Goodbye!!!")
			return
		}
		fmt.Println("Your choice is", choice)
	}

}
func showMenu() (userChoice int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	fmt.Scanln(&userChoice)
	fmt.Println("====================================")
	return userChoice
}
