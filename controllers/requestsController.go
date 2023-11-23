package controllers

import (
	"net/http"
	"time"

	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func CreateRequest(c *gin.Context) {
	var body struct {
		ReqID        int       `json:"reqId"`
		BookID       int       `json:"bookId"`   //fk book
		ReaderID     int       `json:"readerId"` //fk user
		RequestDate  time.Time `json:"requestDate"`
		ApprovalDate time.Time `json:"approvalDate"`
		ApproverID   int       `json:"approverId"` //fk admin
		RequestType  string    `json:"requestType"`
	}

	c.Bind(&body)

	request := models.RequestEvents{
		ReqID:        body.ReqID,
		BookID:       body.BookID,
		ReaderID:     body.ReaderID,
		RequestDate:  body.RequestDate,
		ApprovalDate: body.ApprovalDate,
		ApproverID:   body.ApproverID,
		RequestType:  body.RequestType,
	}

	result := initializers.DB.Create(&request)

	if result.Error != nil{
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book created",
		"request": result,
	})
}

func GetAllRequest(c *gin.Context) {

	var requests []models.RequestEvents

	initializers.DB.Find(&requests)
	c.JSON(http.StatusOK, gin.H{
		"message": "book created",
		"request": requests,
	})


}

func UpdateRequestByReqID(c *gin.Context) {
	
	// get id off the url
	id := c.Param("id")

	// get the data off the req body
	var body struct {
		ReqID        int       `json:"reqId"`
		BookID       int       `json:"bookId"`   //fk book
		ReaderID     int       `json:"readerId"` //fk user
		RequestDate  time.Time `json:"requestDate"`
		ApprovalDate time.Time `json:"approvalDate"`
		ApproverID   int       `json:"approverId"` //fk admin
		RequestType  string    `json:"requestType"`
	}

	c.Bind(&body)

	var request models.RequestEvents

	initializers.DB.Find(&request, id)

	initializers.DB.Model(&request).Updates(models.RequestEvents{
		ReqID:        body.ReqID,
		BookID:       body.BookID,
		ReaderID:     body.ReaderID,
		RequestDate:  body.RequestDate,
		ApprovalDate: body.ApprovalDate,
		ApproverID:   body.ApproverID,
		RequestType:  body.RequestType,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated",
		"request": request,
	})




}
