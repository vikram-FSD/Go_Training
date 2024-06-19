package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func bankDetails() {
	var balance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		return
		// panic("Cant Continue sorry file not found.")
	}
	var withdraw float64 = 0
	var deposits float64 = 0
	for i := 1; i < 200; i++ {
		presentOptions()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("\nAccount_Balance is: %v", balance)
			writeBalanceTOFile("Account_balance is :", balance)
		case 2:
			fmt.Print("enter the amount to withDraw: ")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("WithDrawal amount exceeds the balance, Please enter the correct amount!")
			} else if withdraw <= balance && !(withdraw <= 0) {
				balance -= withdraw
				fmt.Println("New Balance Updated: ", balance)
			} else {
				fmt.Println("Withdraw amount invalid!!!!")
			}
		case 3:
			fmt.Println("enter the amount to be deposited: ")
			fmt.Scan(&deposits)
			if deposits > 0 {
				balance += deposits
				fmt.Println("Balance Updated! New Amount: ", balance)
				writeBalanceTOFile("Deposited Amount: ", deposits)
			} else if deposits <= 0 {
				fmt.Println("Dont enter negative value or Zero, Please Deposit the amount greater than 0 !")
				continue
			}
		case 4:
			fmt.Println("Exited, ThankYou Visit Again..")
			fmt.Println("Reach us 24/7: ", randomdata.FullName(1), randomdata.PhoneNumber())
			return
		default:
			fmt.Println("Enter the valid option !!!")
			return

		}
	}
}

// MULTILINE STRING /WE USE Sprintf() to implement this function.
// multilineString := `This is a multiline string.
// It can span multiple lines
// and contain special characters like "quotes" and backslashes (\).
// You can include tabs or spaces for formatting purposes.`
// var container = fmt.Sprintf("\n MultiLineString: %s", multilineString)
// fmt.Println(container)
// func printnumAndString(num int, str string) {
// 	fmt.Printf("%d\n%s", num, str)
// }
