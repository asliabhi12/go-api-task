package controllers

import (
	"net/http"

	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func CreateLibrary(c *gin.Context) {
	var body struct {
		ID   int    `json:"Id"`
		Name string `json:"name"`
	}

	c.Bind(&body)

	// create a library 

	library := models.Library{Name: body.Name}

	result := initializers.DB.Create(&library) 


	if result.Error != nil {
		c.Status(400)
		return
	}

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
