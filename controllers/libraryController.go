package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLibrary(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message ": "create library",
	})
}

func GetAllLibrary(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message ": "this are all the libraries",
	})
}