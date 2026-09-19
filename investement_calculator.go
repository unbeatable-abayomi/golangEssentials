package main

import (
	"fmt"
	"math"
	"errors"
	"os"
)

const inflationRate = 2.5
const myFinance = "myFinances.txt"
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
     revenue, err := getUserInput("Enter Total Revenue: ")
	 	 if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue, sorry.")
	 }
	 totalExpenses, err := getUserInput("Enter Total Expenses: ")
	 	 if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue, sorry.")
	 }
	 taxRate, err := getUserInput("Enter Tax Rate: ")
	 	 if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue, sorry.")
	 }
     
	// fmt.Print("Enter Total Expenses: ")
	// outputText("Enter Total Expenses: ")
	// fmt.Scan(&totalExpenses)

	// fmt.Print("Enter Tax Rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - totalExpenses

	// profit := ebt * (1 - taxRate/100)

	// ratio := ebt / profit
	// writeBalanceToFile(revenue,totalExpenses,taxRate)
	ebt, profit,ratio := calculateFinacials(revenue,totalExpenses,taxRate)
writeBalanceToFile(ebt,profit,ratio)
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

func outputText(text string)  {
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


func getUserInput(infoText string) (float64,error){
	
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput == 0 {
		return userInput, errors.New("Sorry you entered a Zero Value")
	}
	if userInput < 0{
		return userInput, errors.New("Sorry you entered a Negative Value ")
	}
	return userInput, nil
}


func writeBalanceToFile(ebt,profit,ratio float64){
//   results := fmt.Sprint(revenue,totalExpenses,taxRate)
 results2 := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt,profit,ratio)
  os.WriteFile(myFinance, []byte(results2), 0644)
 
}