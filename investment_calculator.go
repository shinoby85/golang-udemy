package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	//Example
	//var a, b, c, d int = 10, 5.6, "test", true
	//fmt.Println(a, b, c, d)

	var investmentAmount float64 = 1000

	//Example
	//var expectedReturnRate = 5.5

	expectedReturnRate := 5.5
	var years float64 = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)
	//Example
	//var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	//var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjasted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

	//fmt.Printf("Future Value: %.1f\nFuture Value (adjasted for Inflation): %.1f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
