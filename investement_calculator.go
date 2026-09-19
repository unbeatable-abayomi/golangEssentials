package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	fmt.Println("Hello World")

	// const inflationRate = 2.5
	// years, expectedReturnRate := 10.0, 5.5
	var years float64
	var expectedReturnRate float64
	var investmentAmmount float64
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate Amount: ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue:= returnFutureValues(investmentAmmount, expectedReturnRate, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	fmt.Println("Lets calculate EBT, profit and ratio")

	// var revenue float64
	// var totalExpenses float64
	// var taxRate float64

	// fmt.Print("Enter Total Revenue: ")
	// outputText("Enter Total Revenue: ")
	// fmt.Scan(&revenue)
     revenue := getUserInput("Enter Total Revenue: ")
	 totalExpenses := getUserInput("Enter Total Expenses: ")
	 taxRate := getUserInput("Enter Tax Rate: ")
	// fmt.Print("Enter Total Expenses: ")
	// outputText("Enter Total Expenses: ")
	// fmt.Scan(&totalExpenses)

	// fmt.Print("Enter Tax Rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - totalExpenses

	// profit := ebt * (1 - taxRate/100)

	// ratio := ebt / profit
	ebt, profit,ratio := calculateFinacials(revenue,totalExpenses,taxRate)

	fmt.Printf("EBT: %.2f, Profit: %.2f, Ratio: %.2f\n", ebt, profit, ratio)
	fmt.Printf(`
	EBT: %v, 
	Profit: %v, 
	Ratio: %v`,
		ebt, profit, ratio)
	fmt.Printf("EBT: %T, Profit: %T, Ratio: %T\n", ebt, profit, ratio)

	formattedEbt := fmt.Sprintf("Formatted Value for Ebt:  %.1f\n", ebt)
	formattedProfit := fmt.Sprintf("Formatted Value for Profit:  %.1f\n", profit)

	fmt.Print(formattedEbt, formattedProfit)
}

func outputText(text string) {
	fmt.Printf(text)
}

// func returnFutureValues(investmentAmmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }

func returnFutureValues(investmentAmmount, expectedReturnRate, years float64) (fv float64,rfv float64) {
	fv = investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	
}

func calculateFinacials(revenue,totalExpenses,taxRate float64)( float64,float64,float64){
	ebt := revenue - totalExpenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt,profit, ratio

}


func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}