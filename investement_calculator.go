package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	const inflationRate = 2.5
	investmentAmmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	futureValue := investmentAmmount * math.Pow(1 + expectedReturnRate/100, years)
    futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println(futureValue)
    fmt.Println(futureRealValue)

}
