package main

import (
	"fmt"
	"math"
)

const infaltionRate float64 = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue,futureRealValue :=	calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %v\n", futureValue)
	// fmt.Println("Future Value adjusted for inflation: ",futureRealValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)

	formatedFV := fmt.Sprintf("Future Value: %.1f", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f",futureRealValue)
	fmt.Println(formatedFV)
	fmt.Println(formatedRFV)
}


func outputText(text string)  { 
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years);
	rfv = fv / math.Pow(1 + infaltionRate / 100, years);

	return fv, rfv;
}
