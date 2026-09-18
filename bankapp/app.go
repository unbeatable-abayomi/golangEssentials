package main


import (
   "fmt"
)


func main(){
	fmt.Println("Welcome to Bank")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
    fmt.Print("Your Chioce: ")
    var userAction  string
	var initialBalance float64 = 50000
    fmt.Scan(&userAction)

    if userAction == "1"{
		fmt.Printf("Here is the current balance: %v\n", initialBalance)
	} else if userAction == "2"{
       fmt.Print("Pls enter amount to Deposit: ")
	   var despoitAmount  float64
	   fmt.Scan(&despoitAmount)
	   if despoitAmount <= 0 {
		fmt.Println("Hi You entered an invalid amount to deposite")
		return
	   }
	   initialBalance += despoitAmount
	   fmt.Printf("Here is the current balance: %v\n", initialBalance)
	}else if userAction == "3"{
       fmt.Print("Pls enter amount to withdraw: ")
	   
	   var withdramAmount  float64
	   fmt.Scan(&withdramAmount)
	   	if withdramAmount <= 0 || withdramAmount > initialBalance {
		fmt.Println("Hi You entered an invalid amount for withdrawl")
		return
	   }
	   initialBalance -= withdramAmount
	   fmt.Printf("Here is the current balance: %v\n", initialBalance)
	}else{
		println("Goodbye....")
	}

}