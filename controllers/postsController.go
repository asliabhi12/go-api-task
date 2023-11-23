package controllers

import (
	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateBook(c *gin.Context) {
	// get data off req body

	var body struct {
		ISBN            int    `json:"isbn"`
		LibID           int    `json:"libId"`
		Title           string `json:"title"`
		Author          string `json:"author"`
		Publisher       string `json:"publisher"`
		Version         string `json:"version"`
		TotalCopies     int    `json:"totalCopies"`
		AvailableCopies int    `json:"availableCopies"`
	}

	c.Bind(&body)

	// create a book

	book := models.Book{Title: body.Title, ISBN: body.ISBN, LibID: body.LibID, Author: body.Author, Version: body.Version, TotalCopies: body.TotalCopies, AvailableCopies: body.AvailableCopies}

	result := initializers.DB.Create(&book)

	if result.Error != nil {
		c.Status(400)
		return

	}

	c.JSON(200, gin.H{
		"book": book,
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
	var body struct {
		ISBN            int    `json:"isbn"`
		LibID           int    `json:"libId"`
		Title           string `json:"title"`
		Author          string `json:"author"`
		Publisher       string `json:"publisher"`
		Version         string `json:"version"`
		TotalCopies     int    `json:"totalCopies"`
		AvailableCopies int    `json:"availableCopies"`
	}

	c.Bind(&body)

	// find the book we are updating
	var book models.Book
	initializers.DB.Find(&book, id)

	// update it
	initializers.DB.Model(&book).Updates(models.Book{
		Title:           body.Title,
		ISBN:            body.ISBN,
		LibID:           body.LibID,
		Author:          body.Author,
		Version:         body.Version,
		TotalCopies:     body.TotalCopies,
		AvailableCopies: body.AvailableCopies})
	// respond with it

	c.JSON(200, gin.H{
		"books": book,
	})
}
