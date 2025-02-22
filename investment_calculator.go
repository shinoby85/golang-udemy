package main

import (
	"fmt"
	"math"
)

func main() {
	//Example
	//var a, b, c, d int = 10, 5.6, "test", true
	//fmt.Println(a, b, c, d)

	const inflationRate = 2.5
	var investmentAmount float64 = 1000

	//Example
	//var expectedReturnRate = 5.5

	expectedReturnRate := 5.5
	var years float64 = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	//Example
	//var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjasted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

	//fmt.Printf("Future Value: %.1f\nFuture Value (adjasted for Inflation): %.1f\n", futureValue, futureRealValue)
}
