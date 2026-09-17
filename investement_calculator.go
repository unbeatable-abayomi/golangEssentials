package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Println("Hello World")
	   
	const inflationRate = 2.5
	// years, expectedReturnRate := 10.0, 5.5
	// var years float64
	// var expectedReturnRate float64 
	// var investmentAmmount float64   
	// fmt.Print("Investment Amount: ")                                                                                                                                         
	// fmt.Scan(&investmentAmmount)

	// fmt.Print("Years: ")                                                                                                                                         
	// fmt.Scan(&years)

	// fmt.Print("Expected Return Rate Amount: ")                                                                                                                                         
	// fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmmount * math.Pow(1 + expectedReturnRate/100, years)
    // futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// fmt.Println(futureValue)
    // fmt.Println(futureRealValue)

	fmt.Println("Lets calculate EBT, profit and ratio")

	
	var revenue float64
    var totalExpenses float64
	var taxRate  float64

	fmt.Print("Enter Total Revenue: ")
    fmt.Scan(&revenue)

	fmt.Print("Enter Total Expenses: ")
    fmt.Scan(&totalExpenses)

	fmt.Print("Enter Tax Rate: ")
    fmt.Scan(&taxRate)
   
	ebt := revenue - totalExpenses

	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit


	fmt.Printf("EBT: %.2f, Profit: %.2f, Ratio: %.2f\n", ebt,profit,ratio)
	fmt.Printf("EBT: %v, Profit: %v, Ratio: %v\n", ebt,profit,ratio)

}
