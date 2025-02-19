package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	//var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
