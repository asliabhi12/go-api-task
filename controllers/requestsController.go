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

	var Book_requested models.Book

	result_bookFound := initializers.DB.First(&Book_requested, "ISBN = ?", body.BookID)
	if result_bookFound.Error != nil {
		panic(result_bookFound.Error)
	}

	if Book_requested.AvailableCopies > 0 {

		request := models.RequestEvents{
			ReqID:       body.ReqID,
			BookID:      body.BookID,
			ReaderID:    body.ReaderID,
			RequestDate: RequestDate,
			RequestType: body.RequestType,
		}

		result := initializers.DB.Create(&request)

		if result.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "request created"})
	} else {
		// You need to get the nearest date whwn book will be available
		var nearestReturnDate string
		var issueRegistry models.IssueRegistery

		// Query the IssueRegistry table to find the nearest return date for the specified ISBN
		result := initializers.DB.Where("ISBN = ? AND IssueStatus = ?", body.BookID, "Issued").
			Order("ExpectedReturnDate ASC").
			First(&issueRegistry)

		if result.Error != nil {
			nearestReturnDate = "N/A" // Set a default value or handle the case where no records are found
		} else {
			nearestReturnDate = issueRegistry.ExpectedReturnDate
		}
		c.IndentedJSON(400, gin.H{
			"error":              "No available copies for the requested book",
			"nextAvailableDate": nearestReturnDate,
		})
	}

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
		ReqID:       body.ReqID,
		BookID:      body.BookID,
		ReaderID:    body.ReaderID,
		RequestDate: RequestDate,
		RequestType: body.RequestType,
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
