package controllers

import (
	"net/http"

	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func CreateLibrary(c *gin.Context) {
	var body models.Library

	tx := initializers.DB.Begin()
	defer tx.Rollback()

	c.Bind(&body)

	// create a library 

	var library models.Library

	tx.First(&library, "name = ?", body.Name)
	
	if library.Name == body.Name {
		c.JSON(406, gin.H{"message": "Enter a different library name", "Details": body.Name})
		return
	}
	
	library = models.Library{Name: body.Name}
	
	result := tx.Create(&library) 

	if result.Error != nil {
		c.Status(400)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong!!!",
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message ": "created library successfully",
		"library": library,
	})
}

func GetAllLibrary(c *gin.Context) {
	
	// get all the Libraries

	var libraries []models.Library

	initializers.DB.Find(&libraries)

	c.JSON(http.StatusOK, gin.H{
		"message ": "this are all the libraries",
		"libraries": libraries,
	})
}
