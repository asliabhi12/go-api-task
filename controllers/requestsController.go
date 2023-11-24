package controllers

import (
	"net/http"
	"time"
	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func CreateRequest(c *gin.Context) {
	var body models.RequestEvents
	c.Bind(&body)
	now := time.Now()
	RequestDate := now.Format("2006-01-02")



	request := models.RequestEvents{
		ReqID:        body.ReqID,
		BookID:       body.BookID,
		ReaderID:     body.ReaderID,
		RequestDate:  RequestDate,
		RequestType:  body.RequestType,
	}

	result := initializers.DB.Create(&request)

	if result.Error != nil{
		c.Status(400)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "request created",
	// 	"request": result,
	// })

	c.JSON(http.StatusOK, gin.H{"message": "request created"})
}

func GetAllRequest(c *gin.Context) {

	var requests []models.RequestEvents

	initializers.DB.Find(&requests)
	c.JSON(http.StatusOK, gin.H{
		"message": "Got all requests",
		"request": requests,
	})


}

func UpdateRequestByReqID(c *gin.Context) {
	
	// get id off the url
	id := c.Param("id")

	// get the data off the req body
	var body models.RequestEvents

	c.Bind(&body)

	var request models.RequestEvents

	initializers.DB.Find(&request, id)

	now := time.Now()
	RequestDate := now.Format("2006-01-02")


	initializers.DB.Model(&request).Updates(models.RequestEvents{
		ReqID:        body.ReqID,
		BookID:       body.BookID,
		ReaderID:     body.ReaderID,
		RequestDate:  RequestDate,
		RequestType:  body.RequestType,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated",
		"request": request,
	})

}


func GetRequest(c *gin.Context) {
	// get id off the url
	id := c.Param("id")

	// get the Request

	var request models.RequestEvents
	initializers.DB.Find(&request, id)

	// resoponding with them

	c.JSON(200, gin.H{
		"request": request,
	})

}