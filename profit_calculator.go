package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println()
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeCalculateDataToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("Revenue is zero or negative")
		return 0, errors.New("Value must be a positive number.")
	}
	return userInput, nil
}

func writeCalculateDataToFile(ebt, profit, ratio float64) {
	calcData := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.3f\n", ebt, profit, ratio)
	os.WriteFile("calc.txt", []byte(calcData), 0644)
}
