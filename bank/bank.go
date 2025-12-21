package bank

import "fmt"

type BankAccount struct {
	Balance int
}

func (account *BankAccount) Deposit(amount int) {
	account.Balance += amount
}
func (accoumt *BankAccount) Withdraw(amount int) {
	if amount > accoumt.Balance {
		fmt.Println("Not enough balance")
		return
	}
	accoumt.Balance -= amount
}
func (account *BankAccount) GetBalance() int {
	return account.Balance
}
func BankCLI() {
	fmt.Println("Welcome to bank Account Simulation application!")

	for loop := true; loop; {

		var choice uint8

		fmt.Println("==============================")
		fmt.Println("Deposit: 1")
		fmt.Println("Withdraw: 2")
		fmt.Println("Get Balance: 3")
		fmt.Println("Exit: 4")
		fmt.Println("==============================")

		fmt.Println("What to do now?")
		fmt.Scan(&choice)

		switch choice {
		case 1:
		case 2:
		case 3:
		case 4:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
