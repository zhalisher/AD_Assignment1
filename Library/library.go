package Library

import (
	"fmt"
)

type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

//type Library struct {
//	Book map[string]Book
//}

func library() {
	fmt.Println("Welcome to Library Management System application!")

	for loop := true; loop; {

		var choice uint8

		fmt.Println("==============================")
		fmt.Println("Add book: 1")
		fmt.Println("Borrow book: 2")
		fmt.Println("ReturnBook: 3")
		fmt.Println("List all available book: 4")
		fmt.Println("Exit: 5")
		fmt.Println("==============================")

		fmt.Println("What to do now?")
		fmt.Scan(&choice)

		switch choice {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
