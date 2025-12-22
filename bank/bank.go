package bank

import "fmt"

type BankAccount struct {
	Balance int
}

func (account *BankAccount) Deposit(amount int) {
	if amount > 0 {
		account.Balance += amount
	}
}
func (accoumt *BankAccount) Withdraw(amount int) {
	if amount > accoumt.Balance {
		fmt.Println("Not enough balance")
		return
	}
	accoumt.Balance -= amount
}

type Transction struct {
	Action string
	Amount int
}

func (account *BankAccount) GetBalance() int {
	return account.Balance
}
func BankCLI() {
	account := &BankAccount{}
	var transactions []Transction

	fmt.Println("Welcome to bank Account Simulation application!")

	for loop := true; loop; {

		var choice uint8

		fmt.Println("==============================")
		fmt.Println("Deposit: 1")
		fmt.Println("Withdraw: 2")
		fmt.Println("Get Balance: 3")
		fmt.Println("Show transactions: 4")
		fmt.Println("Exit: 5")
		fmt.Println("==============================")

		fmt.Println("What to do now?")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var amount int
			println("Enter amount of money to deposit")
			fmt.Scan(&amount)
			account.Deposit(amount)
			transactions = append(transactions, Transction{"Deposit", amount})
		case 2:
			var amount int
			fmt.Println("Enter amount of money to withdraw")
			fmt.Scan(&amount)
			account.Withdraw(amount)
			transactions = append(transactions, Transction{"Withdraw", amount})
		case 3:
			fmt.Println(account.GetBalance())
		case 4:
			fmt.Println("Transactions::")
			for i, j := range transactions {
				fmt.Printf("No: %v. Action: %v. Amount: %v\n", i+1, j.Action, j.Amount)
			}
		case 5:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
