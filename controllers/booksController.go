package controllers

import (
	"fmt"

	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	// get data off req body

	var body models.Book
	c.Bind(&body)

	tx := initializers.DB.Begin()
	defer tx.Rollback()
	// create a book

	var existingBook models.Book
	result := tx.First(&existingBook, "ISBN = ?", body.ISBN)

	if result.Error == nil {
		// Book already exists, increment the number of available copies
		existingBook.AvailableCopies += body.AvailableCopies
		result = tx.Save(&existingBook)

		if result.Error != nil {
			fmt.Println("Error during save:", result.Error)
			tx.Rollback()
			c.Status(400)
			return
		}

		// Commit the transaction if everything is successful
		tx.Commit()

		c.JSON(200, gin.H{
			"book":    existingBook,
			"message": "Book already exists. Available copies updated successfully",
		})
		return
	}

	book := models.Book{
		Title:           body.Title,
		ISBN:            body.ISBN,
		LibID:           body.LibID,
		Author:          body.Author,
		Version:         body.Version,
		TotalCopies:     body.TotalCopies,
		AvailableCopies: body.AvailableCopies,
		Publisher:       body.Publisher,
	}

	result = tx.Create(&book)

	if result.Error != nil {
		c.Status(400)
		return

	}
	tx.Commit()
	c.JSON(200, gin.H{
		"book":    book,
		"message": "Book created successfully",
	})
}

func BooksIndex(c *gin.Context) {
	// get the books
	var books []models.Book
	initializers.DB.Find(&books)

	c.JSON(200, gin.H{
		"books": books,
	})

}

func BookShow(c *gin.Context) {
	// get id off the url
	id := c.Param("id")

	// get the books

	var book models.Book
	initializers.DB.Find(&book, id)

	// resoponding with them

	c.JSON(200, gin.H{
		"books": book,
	})

}

func BooksUpdate(c *gin.Context) {
	// get id off the url
	id := c.Param("id")

	// get the data off the req body
	var body models.Book

	c.Bind(&body)

	// find the book we are updating
	var book models.Book
	initializers.DB.Find(&book, id)

	// update it
	initializers.DB.Model(&book).Updates(models.Book{
		Title:           body.Title,
		ISBN:            body.ISBN,
		LibID:           body.LibID,
		Publisher:       body.Publisher,
		Author:          body.Author,
		Version:         body.Version,
		TotalCopies:     body.TotalCopies,
		AvailableCopies: body.AvailableCopies})
	// respond with it

	c.JSON(200, gin.H{
		"books": book,
	})
}

func DeleteBook(c *gin.Context) {
	// Get book ID from the URL parameter
	id := c.Param("id")

	// Check if the book with the specified ID exists
	var existingBook models.Book
	if result := initializers.DB.First(&existingBook, id); result.Error != nil {
		// If the book is not found, respond with a 404 status code
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	// Delete the book
	if result := initializers.DB.Delete(&existingBook); result.Error != nil {
		// If an error occurs during deletion, respond with a 500 status code and an error message
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	// Respond with a success message
	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
	})
}
