package services

import (
	"afl3/models"
	"errors"
	"strconv"
)

var books = []models.Book{
	{ID: "1", Title: "Belajar Membaca", Author: "Nizam", Price: 9.99},
	{ID: "2", Title: "Kisah Cinta", Author: "Raya", Price: 14.99},
	{ID: "3", Title: "Pedoman Hidup", Author: "Tamtama", Price: 12.99},
	{ID: "4", Title: "Belajar Adil", Author: "Agung", Price: 19.99},
}

// GetAllBooks returns all books
func GetAllBooks() []models.Book {
	return books
}

// GetBookByID returns a book by its ID
func GetBookByID(id string) (models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func AddBook(book models.Book) {
	// Ambil ID terbesar dari buku yang ada
	maxID := 0
	for _, b := range books {
		idInt, _ := strconv.Atoi(b.ID)
		if idInt > maxID {
			maxID = idInt
		}
	}

	// Set ID baru sebagai maxID + 1
	book.ID = strconv.Itoa(maxID + 1)

	// Tambahkan buku baru ke daftar
	books = append(books, book)
}

// UpdateBook updates an existing book
func UpdateBook(id string, updatedBook models.Book) error {
	for i, book := range books {
		if book.ID == id {
			// Memperbarui hanya field yang relevan
			books[i].Title = updatedBook.Title
			books[i].Author = updatedBook.Author
			books[i].Price = updatedBook.Price
			return nil
		}
	}
	return errors.New("book not found")
}

// DeleteBook deletes a book by its ID
func DeleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
