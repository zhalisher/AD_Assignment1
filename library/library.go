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
	library.Books[book.Title] = book
	fmt.Println("Book added: ", book.Title)
}
func (library *Library) BorrowBook(id int) {

}
func (library *Library) ReturnBook(id int) {

}

func (library Library) ListAvailableBooks() {
	fmt.Println("Available books:")
	for _, books := range library.Books {
		fmt.Printf("Id: %v, Title: %v, Author: %v\n", books.ID, books.Title, books.Author)
	}
}

func LibraryCLI() {
	library := Library{Books: make(map[string]Book)}
	id := 0
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
			var title, author string

			fmt.Println("Etner book title: ")
			fmt.Scan(&title)

			fmt.Println("Enter book author: ")
			fmt.Scan(&author)

			id++

			library.AddBook(Book{ID: id, Title: title, Author: author})
		case 2:
		case 3:
		case 4:
			library.ListAvailableBooks()
		case 5:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
