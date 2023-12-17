package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
)

func ApproveRequest(c *gin.Context) {
	var body models.RequestEvents
	c.Bind(&body)
	now := time.Now()
	approverID := c.MustGet("approverId").(int)
	body.ApproverID = approverID
	fmt.Println("********************apprid", body.ApproverID)
	fmt.Println("********************reqid", body.ReqID)

	approvalDate := now.Format("2006-01-02")

	body.ApprovalDate = approvalDate

	tx := initializers.DB.Begin()
	defer tx.Rollback()

	var reqData models.RequestEvents
	// if err := tx.First(&reqData, "req_id = ?", body.ReqID).Error; err == nil {
	// 	tx.Rollback()
	// 	c.JSON(400, gin.H{"error": "Request Not Found"})
	// 	return
	// }
	result := tx.First(&reqData, "req_id = ?", body.ReqID)
	if result.RowsAffected <= 0 {
		fmt.Println(result.RowsAffected)
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Request Not Found"})
		return
	}

	// Check if the request already exists in the issue_registry table
	var existingIssue models.IssueRegistery
	if err := tx.First(&existingIssue, "isbn = ? AND reader_id = ?", reqData.BookID, reqData.ReaderID).Error; err == nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Request already approved and book issued"})
		return
	}

	// Check book availability in the books table
	var book models.Book
	if err := tx.First(&book, "isbn = ?", reqData.BookID).Error; err != nil {
		tx.Rollback()
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies <= 0 {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "No available copies for the book"})
		return
	}

	// Set the request with the approver ID and approval date
	body.ApprovalDate = approvalDate
	body.ApproverID = approverID
	body.RequestType = "issued"
	// Update the request in the database

	var updateRequest models.RequestEvents

	tx.Find(&updateRequest, reqData.ReqID)

	// tx.Model(&updateRequest).Updates(models.RequestEvents{
	// 	ApprovalDate: body.ApprovalDate,
	// 	ApproverID:   body.ApproverID,
	// })

	var temp models.RequestEvents
	tx.Model(&temp).Where("req_id = ?", reqData.ReqID).Updates(body)
	fmt.Println("**************************", body)

	// // Create a new entry in the Issue Registry
	issueRegistry := models.IssueRegistery{
		ISBN:               reqData.BookID,
		ReaderID:           reqData.ReaderID,
		IssueApproverID:    approverID,
		IssueStatus:        "Issued",
		IssueDate:          approvalDate,
		ExpectedReturnDate: now.AddDate(0, 0, 15).Format("2006-01-02"), // Adding 15 days to IssueDate
	}
	// Insert the new entry into the Issue Registry
	if err := tx.Create(&issueRegistry).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to create entry in Issue Registry"})
		return
	}

	// Update the available book count by decreasing it
	if err := tx.Model(&book).Update("available_copies", book.AvailableCopies-1).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to update book availability"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Request approved successfully"})
}


func DenyRequest(c *gin.Context) {
	var body models.RequestEvents
	c.Bind(&body)
	now := time.Now()
	approverID := c.MustGet("approverId").(int)
	body.ApproverID = approverID
	fmt.Println("********************apprid", body.ApproverID)
	fmt.Println("********************reqid", body.ReqID)

	approvalDate := now.Format("2006-01-02")

	body.ApprovalDate = approvalDate

	tx := initializers.DB.Begin()
	defer tx.Rollback()

	var reqData models.RequestEvents
	// if err := tx.First(&reqData, "req_id = ?", body.ReqID).Error; err == nil {
	// 	tx.Rollback()
	// 	c.JSON(400, gin.H{"error": "Request Not Found"})
	// 	return
	// }
	result := tx.First(&reqData, "req_id = ?", body.ReqID)
	if result.RowsAffected <= 0 {
		fmt.Println(result.RowsAffected)
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Request Not Found"})
		return
	}

	// Check if the request already exists in the issue_registry table
	var existingIssue models.IssueRegistery
	if err := tx.First(&existingIssue, "isbn = ? AND reader_id = ?", reqData.BookID, reqData.ReaderID).Error; err == nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Request already Denied and book not issued"})
		return
	}

	// Check book availability in the books table
	var book models.Book
	if err := tx.First(&book, "isbn = ?", reqData.BookID).Error; err != nil {
		tx.Rollback()
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies <= 0 {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "No available copies for the book"})
		return
	}

	// Set the request with the approver ID and approval date
	body.ApprovalDate = approvalDate
	body.ApproverID = approverID
	body.RequestType = "rejected"
	// Update the request in the database

	var updateRequest models.RequestEvents

	tx.Find(&updateRequest, reqData.ReqID)

	// tx.Model(&updateRequest).Updates(models.RequestEvents{
	// 	ApprovalDate: body.ApprovalDate,
	// 	ApproverID:   body.ApproverID,
	// })

	var temp models.RequestEvents
	tx.Model(&temp).Where("req_id = ?", reqData.ReqID).Updates(body)
	fmt.Println("**************************", body)

	// // Create a new entry in the Issue Registry
	issueRegistry := models.IssueRegistery{
		ISBN:               reqData.BookID,
		ReaderID:           reqData.ReaderID,
		IssueApproverID:    approverID,
		IssueStatus:        "Rejected",
		IssueDate:          approvalDate,
		
	}
	// Insert the new entry into the Issue Registry
	if err := tx.Create(&issueRegistry).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to create entry in Issue Registry"})
		return
	}

	// Update the available book count by decreasing it
	if err := tx.Model(&book).Update("available_copies", book.AvailableCopies).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to update book availability"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Request Denied successfully"})
}



// get information about issued



func IssuesIndex(c *gin.Context) {
	// get the books
	var issues []models.IssueRegistery
	initializers.DB.Find(&issues)

	c.JSON(200, gin.H{
		"issued": issues,
	})

}