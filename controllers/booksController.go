package controllers

import (
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

	book := models.Book{
		Title: body.Title,
		ISBN: body.ISBN,
		LibID: body.LibID, 
		Author: body.Author, 
		Version: body.Version, 
		TotalCopies: body.TotalCopies, 
		AvailableCopies: body.AvailableCopies,
		Publisher: body.Publisher,
	}

	result := tx.Create(&book)

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
