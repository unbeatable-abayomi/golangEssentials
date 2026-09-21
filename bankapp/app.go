package main


import (
   "fmt"
   "os"
   "strconv"
   "errors"
)

const accountBalanceFIle = "balance.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFIle)
	if err != nil {
		return 1000, errors.New("Falied To find Balance File")
	}
	balanceText := string(data)
	balance, err:= strconv.ParseFloat(balanceText, 64)

		if err != nil {
		return 1000, errors.New("Falied To Parse Balance File")
	}

	return balance, nil

}
func writeBalanceToFile(balance float64){
  balanceText := fmt.Sprint(balance)
  os.WriteFile(accountBalanceFIle, []byte(balanceText), 0644)

}


func main(){

//var initialBalance float64 = 
    initialBalance, err := getBalanceFromFile()

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue, sorry.")
	}
 for {
	// fmt.Println("Welcome to Bank")
	// fmt.Println("What do you want to do")
	// fmt.Println("1. Check Balance")
	// fmt.Println("2. Deposit")
	// fmt.Println("3. Withdraw")
	// fmt.Println("4. Exit")
    // fmt.Print("Your Chioce: ")
	presentOptions()
    var userAction  string
	
    fmt.Scan(&userAction)
    switch userAction{
	case "1":
				fmt.Printf("Here is the current balance: %v\n", initialBalance)
	case "2":
		       fmt.Print("Pls enter amount to Deposit: ")
	   var despoitAmount  float64
	   fmt.Scan(&despoitAmount)
	   if despoitAmount <= 0 {
		fmt.Println("Hi You entered an invalid amount to deposite")
		//return
		continue
	   }
	   initialBalance += despoitAmount
	   fmt.Printf("Here is the current balance: %v\n", initialBalance)
	   writeBalanceToFile(initialBalance)
	case "3":
		fmt.Print("Pls enter amount to withdraw: ")
	   
	   var withdramAmount  float64
	   fmt.Scan(&withdramAmount)
	   	if withdramAmount <= 0 {
		fmt.Println("Hi You entered an invalid amount to withdraw below 0 or 0")
		//return
		continue

	   }

	   if withdramAmount > initialBalance {
	fmt.Println("Hi You entered an invalid amount to withdraw above balance")
		//return
		continue
	   }
	   initialBalance -= withdramAmount
	   fmt.Printf("Here is the current balance: %v\n", initialBalance)
	   writeBalanceToFile(initialBalance)
	default:
		fmt.Println("Goodbye....")
		fmt.Println("Thank you for Using our Bank")
		return
		//break

	}
    // if userAction == "1"{
	// 	fmt.Printf("Here is the current balance: %v\n", initialBalance)
	// } else if userAction == "2"{
    //    fmt.Print("Pls enter amount to Deposit: ")
	//    var despoitAmount  float64
	//    fmt.Scan(&despoitAmount)
	//    if despoitAmount <= 0 {
	// 	fmt.Println("Hi You entered an invalid amount to deposite")
	// 	//return
	// 	continue
	//    }
	//    initialBalance += despoitAmount
	//    fmt.Printf("Here is the current balance: %v\n", initialBalance)
	// }else if userAction == "3"{
    //    fmt.Print("Pls enter amount to withdraw: ")
	   
	//    var withdramAmount  float64
	//    fmt.Scan(&withdramAmount)
	//    	if withdramAmount <= 0 {
	// 	fmt.Println("Hi You entered an invalid amount to withdraw below 0 or 0")
	// 	//return
	// 	continue

	//    }

	//    if withdramAmount > initialBalance {
	// fmt.Println("Hi You entered an invalid amount to withdraw above balance")
	// 	//return
	// 	continue
	//    }
	//    initialBalance -= withdramAmount
	//    fmt.Printf("Here is the current balance: %v\n", initialBalance)
	// }else{
	// 	println("Goodbye....")
	// 	//return
	// 	break
	// }
	}
  //fmt.Println("Thank you for Using our Bank")
}

// func presentOptions2(){
// 		fmt.Println("Welcome to Bank")
// 	fmt.Println("What do you want to do")
// 	fmt.Println("1. Check Balance")
// 	fmt.Println("2. Deposit")
// 	fmt.Println("3. Withdraw")
// 	fmt.Println("4. Exit")
//     fmt.Print("Your Chioce: ")
// }