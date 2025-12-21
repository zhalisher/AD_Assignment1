package library

import (
	"fmt"
)

type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (library *Library) AddBook(book Book) {
	library.Books = make(map[string]Book)
	library.Books[book.Title] = book
	fmt.Println("Book added: ", book.Title)
}
func (library *Library) BorrowBook(book Book) {
	//isExist := library.Books[book]

}
func (library *Library) ReturnBook(book Book) {

}

func LibraryCLI() {
	library := Library{Books: make(map[string]Book)}
	fmt.Println("Welcome to library Management System application!")

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
			var id int
			var title, author string

			fmt.Println("Enter book id: ")
			fmt.Scan(&id)

			fmt.Println("Etner book title: ")
			fmt.Scan(&title)

			fmt.Println("Enter book autor: ")
			fmt.Scan(&author)

			library.AddBook(Book{ID: id, Title: title, Author: author})
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
