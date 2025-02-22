package main

import "fmt"

func main() {
	accountBalance := 10000.0
	choice := showMenu()

	if choice == 1 {
		fmt.Printf("Your balance is %.2f\n", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
		accountBalance += deposit
		fmt.Println("Your balance updated! New amount: ", accountBalance)
	}
	fmt.Println("Your choice is", choice)
}
func showMenu() (userChoice int) {
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	fmt.Scanln(&userChoice)
	fmt.Println("====================================")
	return
}
