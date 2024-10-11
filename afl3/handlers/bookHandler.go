// handlers/Handler.go
package handlers

import (
	"afl3/models"
	"afl3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks - Get all books
func GetBooks(c *gin.Context) {
	books := services.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

// GetBookByID - Get a specific book by ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook - Add a new book
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	services.AddBook(newBook)
	c.JSON(http.StatusCreated, newBook)
}

// UpdateBook - Update an existing book
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := services.UpdateBook(id, updatedBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook - Delete a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
