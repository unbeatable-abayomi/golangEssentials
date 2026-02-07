package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	const inflationRate = 2.5
	years, expectedReturnRate := 10.0, 5.5
	var investmentAmmount float64                                                                                                                                            
	fmt.Scan(&investmentAmmount)
	futureValue := investmentAmmount * math.Pow(1 + expectedReturnRate/100, years)
    futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println(futureValue)
    fmt.Println(futureRealValue)

}
