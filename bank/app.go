package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const filename = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(filename)

	if err != nil {
		fmt.Println("Error reading file, setting balance to 1000")
		accountBalance = 1000
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter the amount you want to deposit: ")
			var depositedAmount float64
			fmt.Scan(&depositedAmount)
			accountBalance += depositedAmount
			fmt.Println("Your account balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, filename)
		case 3:
			fmt.Print("Enter the amount you want to withdraw: ")
			var amountRemoved float64
			fmt.Scan(&amountRemoved)
			accountBalance -= amountRemoved
			fmt.Println("Your account balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, filename)
		default:
			fmt.Println("Goodbye!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your account balance is: ", accountBalance)
		// 	continue
		// } else if choice == 2 {
		// 	fmt.Print("Enter the amount you want to deposit: ")
		// 	var depositedAmount float64
		// 	fmt.Scan(&depositedAmount)
		// 	accountBalance += depositedAmount
		// 	fmt.Println("Your account balance is: ", accountBalance)
		// 	continue
		// } else if choice == 3 {
		// 	fmt.Print("Enter the amount you want to withdraw: ")
		// 	var amountRemoved float64
		// 	fmt.Scan(&amountRemoved)
		// 	accountBalance -= amountRemoved
		// 	fmt.Println("Your account balance is: ", accountBalance)
		// 	continue
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}
}
