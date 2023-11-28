package controllers

import (
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
	approvalDate := now.Format("2006-01-02")

	tx := initializers.DB.Begin()
	defer tx.Rollback()

	// Check if the request already exists in the issue_registry table
	var existingIssue models.IssueRegistery
	if err := tx.First(&existingIssue, "isbn = ? AND reader_id = ?", body.BookID, body.ReaderID).Error; err == nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "Request already approved and book issued"})
		return
	}


	// Check book availability in the books table
	var book models.Book
	if err := tx.First(&book, "isbn = ?", body.BookID).Error; err != nil {
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

	// Update the request in the database


	var updateRequest models.RequestEvents

	tx.Find(&updateRequest, body.ReqID)


	tx.Model(&updateRequest).Updates(models.RequestEvents{
		ApprovalDate: body.ApprovalDate,
		ApproverID: body.ApproverID,
	})

	



	// Create a new entry in the Issue Registry
	issueRegistry := models.IssueRegistery{
		ISBN:               body.BookID,
		ReaderID:           body.ReaderID,
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

