package main

import "fmt"

var (
	balance  int64 = 12000
	withdraw int64 = 0
	deposits int64 = 0
)

func bankDetails() {
	for i := 1; i < 200; i++ {
		fmt.Println("\nEnter your choice: ")
		fmt.Println("1. Check_Account_Balance")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Deposit Money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("\nAccount_Balance is: %v", balance)
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
			} else if deposits <= 0 {
				fmt.Println("Dont enter negative value or Zero, Please Deposit the amount greater than 0 !")
				continue
			}
		case 4:
			fmt.Println("Exited, ThankYou Visit Again..")
			return
		default:
			fmt.Println("Enter the valid option !!!")
			return

		}
	}
}

//MULTILINE STRING /WE USE Sprintf() to implement this function.
// multilineString := `This is a multiline string.
// It can span multiple lines
// and contain special characters like "quotes" and backslashes (\).
// You can include tabs or spaces for formatting purposes.`
// var container = fmt.Sprintf("\n MultiLineString: %s", multilineString)
// fmt.Println(container)
// func printnumAndString(num int, str string) {
// 	fmt.Printf("%d\n%s", num, str)
// }
