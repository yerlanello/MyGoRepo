package library

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists && !book.IsBorrowed {
		book.IsBorrowed = true
		l.Books[id] = book
		fmt.Println("Book borrowed successfully.")
	} else {
		fmt.Println("Book is not available.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.Books[id] = book
		fmt.Println("Book returned successfully.")
	} else {
		fmt.Println("Book was not borrowed.")
	}
}

func (l *Library) ListBooks() {
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
