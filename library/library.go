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
func (library *Library) BorrowBook(title string) {
	book := library.Books[title]
	book.IsBorrowed = true
	library.Books[book.Title] = book
}
func (library *Library) ReturnBook(title string) {
	book := library.Books[title]
	book.IsBorrowed = false
	library.Books[title] = book
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
		fmt.Println("List all borrowed book: 5")
		fmt.Println("Exit: 6")
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
			var title string

			fmt.Println("Enter book title to borrow")
			fmt.Scan(&title)
			book, ok := library.Books[title]
			if !ok {
				fmt.Println("Book not found")
			} else if book.IsBorrowed {
				fmt.Println("Book is already borrowed")
			} else {
				library.BorrowBook(title)
				fmt.Println("You borrowed:", title)
			}
		case 3:
			var title string

			fmt.Println("Enter book title to return")
			fmt.Scan(&title)
			book, ok := library.Books[title]
			if !ok {
				fmt.Println("Book not found")
			} else if !book.IsBorrowed {
				fmt.Println("Book was not borrowed")
			} else {
				library.ReturnBook(title)
				fmt.Println("You returned:", title)
			}
		case 4:
			library.ListAvailableBooks()
		case 5:
			fmt.Println("Borrowed books: ")
			hasBorrowed := false
			for _, b := range library.Books {
				if b.IsBorrowed {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
					hasBorrowed = true
				}
			}
			if !hasBorrowed {
				fmt.Println("No borrowed books")
			}

		case 6:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
