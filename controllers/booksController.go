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


// func RemoveBook(c *gin.Context) {
// 	// Retrieve ISBN from the request URL parameters
// 	isbn := c.Param("isbn")

// 	tx := initializers.DB.Begin()
// 	defer tx.Rollback()

// 	// Check if the book exists based on ISBN
// 	var existingBook models.Book
// 	result := tx.First(&existingBook, "ISBN = ?", isbn)

// 	if result.Error != nil {
// 		c.Status(404) // Book not found
// 		return
// 	}

// 	// Check if there are any issued copies of the book
// 	var issuedCopies int64
// 	tx.Model(&models.IssueRegistery{}).Where("ISBN = ? AND IssueStatus = ?", isbn, "Issued").Count(&issuedCopies)

// 	if issuedCopies > 0 {
// 		c.IndentedJSON(400, gin.H{
// 			"error":   "Cannot remove book with issued copies",
// 			"message": "Please return all issued copies before removing the book",
// 		})
// 		return
// 	}

// 	// Decrement the available copies
// 	existingBook.AvailableCopies -= existingBook.AvailableCopies // Adjust this based on your specific logic

// 	// Save the changes to the existing book
// 	result = tx.Save(&existingBook)

// 	if result.Error != nil {
// 		fmt.Println("Error during save:", result.Error)
// 		tx.Rollback()
// 		c.Status(400)
// 		return
// 	}

// 	// Commit the transaction if everything is successful
// 	tx.Commit()

// 	c.JSON(200, gin.H{
// 		"book":    existingBook,
// 		"message": "Book removed successfully",
// 	})
// }


package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/models" // Update with your actual package path
	"github.com/yourusername/yourproject/initializers" // Update with your actual package path
)

// DeleteBook deletes a book based on ISBN
func DeleteBook(c *gin.Context) {
	// Retrieve ISBN from the URL parameters
	isbn := c.Param("isbn")

	// Begin a transaction
	tx := initializers.DB.Begin()
	defer tx.Rollback()

	// Check if the book exists based on ISBN
	var existingBook models.Book
	result := tx.First(&existingBook, "ISBN = ?", isbn)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with ISBN %s not found", isbn),
		})
		return
	}

	// Check if there are any issued copies of the book
	var issuedCopies int64
	tx.Model(&models.IssueRegistery{}).Where("ISBN = ? AND IssueStatus = ?", isbn, "Issued").Count(&issuedCopies)

	if issuedCopies > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot delete book with issued copies",
			"message": "Please return all issued copies before deleting the book",
		})
		return
	}

	// Delete the book
	result = tx.Delete(&existingBook)

	if result.Error != nil {
		fmt.Println("Error during delete:", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Error deleting the book",
		})
		return
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ISBN %s deleted successfully", isbn),
	})
}
