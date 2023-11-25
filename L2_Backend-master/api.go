package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Server Started"})
}

// ******************
// Admin API's
// ******************

func createLibrary(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	Name := c.Request.FormValue("name")

	// Save file metadata to database
	finalData := Library{
		Name: Name,
	}

	//DATABASE OPERATIONS
	DB := db_connection()

	tx := DB.Begin()
	defer tx.Rollback()

	var library Library

	tx.First(&library, "name = ?", finalData.Name)

	if library.Name == finalData.Name {
		c.JSON(406, gin.H{"message": "Enter a different library name", "Details": finalData})
	} else {
		// Create
		result := tx.Create(&finalData)

		if result.Error != nil {
			panic(result.Error)
		}
		// print(result.RowsAffected)

		var user Users
		tx.First(&library, "name = ?", finalData.Name)
		user.Role = "owner"
		user.LibID = library.ID
		result1 := tx.Create(&user)
		if result1.Error != nil {
			panic(result1.Error)
		}
		tx.Commit()
		c.JSON(201, gin.H{"message": "Library added successfully", "Details": finalData})
	}
}

func createUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	Name := c.Request.FormValue("name")
	Email := c.Request.FormValue("email")
	ContactNumber := c.Request.FormValue("contact")
	LibID := c.Request.FormValue("libId")
	Password := c.Request.FormValue("password")

	u64_libId, err := strconv.ParseUint(LibID, 10, 32)
	if err != nil {
		panic(err)
	}
	libId := uint(u64_libId)

	// Save file metadata to database
	finalData := Users{
		Name:          Name,
		Email:         Email,
		ContactNumber: ContactNumber,
		LibID:         libId,
		Password:      Password,
	}

	//DATABASE OPERATIONS

	DB := db_connection()

	var library Library

	DB.First(&library, "ID = ?", finalData.LibID)

	if library.ID == finalData.LibID {
		var usr Users
		result_checkRecord := DB.First(&usr, "role = ?", "admin")

		if result_checkRecord.RowsAffected > 0 {
			if usr.Role != "admin" {
				finalData.Role = "admin"
				result1 := DB.Create(&finalData)

				if result1.Error != nil {
					panic(result1.Error)
				}
				// print(result.RowsAffected)
			} else {
				finalData.Role = "reader"
				result1 := DB.Create(&finalData)

				if result1.Error != nil {
					panic(result1.Error)
				}
				// print(result.RowsAffected)
			}
		} else {
			finalData.Role = "admin"
			result1 := DB.Create(&finalData)

			if result1.Error != nil {
				panic(result1.Error)
			}
			// print(result.RowsAffected)
		}
		c.JSON(201, gin.H{"message": "User added successfully", "Details": finalData})
	} else {
		c.JSON(406, gin.H{"message": "LibId Do not Match in Library, Enter a valid one", "Details": finalData})
	}
}

func login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	Email := c.Request.FormValue("email")
	Password := c.Request.FormValue("password")

	//DATABASE OPERATIONS

	DB := db_connection()

	var usr Users

	result_findUser := DB.First(&usr, "email = ?", Email)

	if result_findUser.RowsAffected == 0 {
		c.JSON(404, gin.H{"message": "User Not Found"})
	} else {
		if usr.Password == Password {
			c.JSON(201, gin.H{"message": "User Logged In successfully"})
		} else {
			c.JSON(400, gin.H{"message": "User Password Incorrect"})
		}
	}
}

func addBooks(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	ISBN := c.Request.FormValue("isbn")
	LibID := c.Request.FormValue("libId")
	Title := c.Request.FormValue("title")
	Authors := c.Request.FormValue("authors")
	Publisher := c.Request.FormValue("publisher")
	Version := c.Request.FormValue("version")
	TotalCopies := c.Request.FormValue("totalCopies")
	AdminEmail := c.Request.FormValue("adminEmail")

	u64_Id, err := strconv.ParseUint(ISBN, 10, 32)
	if err != nil {
		panic(err)
	}
	isbn := uint(u64_Id)

	u64_libId, err := strconv.ParseUint(LibID, 10, 32)
	if err != nil {
		panic(err)
	}
	libId := uint(u64_libId)

	u64_version, err := strconv.ParseUint(Version, 10, 32)
	if err != nil {
		panic(err)
	}
	version := uint(u64_version)

	u64_totalCopies, err := strconv.ParseUint(TotalCopies, 10, 32)
	if err != nil {
		panic(err)
	}
	totalCopies := uint(u64_totalCopies)

	finalData := BookInventory{
		ISBN:            isbn,
		LibID:           libId,
		Title:           Title,
		Authors:         Authors,
		Publisher:       Publisher,
		Version:         version,
		TotalCopies:     totalCopies,
		AvailableCopies: totalCopies,
	}

	//DATABASE OPERATIONS

	DB := db_connection()

	var Book BookInventory

	var usr Users
	DB.First(&usr, "email = ?", AdminEmail)
	if usr.Role != "admin" {
		c.JSON(400, gin.H{"message": "Unknown User", "Details": finalData})
		return
	}

	DB.First(&Book, "ISBN = ?", finalData.ISBN)

	// When no previous book is there
	if Book.ISBN != finalData.ISBN {
		result1 := DB.Create(&finalData)
		if result1.Error != nil {
			panic(result1.Error)
		}
		// print(result.RowsAffected)

		c.JSON(201, gin.H{"message": "Book added successfully", "Details": finalData})
	} else {
		finalData.TotalCopies = Book.TotalCopies + finalData.TotalCopies
		finalData.AvailableCopies = Book.AvailableCopies + finalData.AvailableCopies

		DB.Model(&Book).Where("ISBN = ?", finalData.ISBN).Updates(BookInventory{TotalCopies: finalData.TotalCopies, AvailableCopies: finalData.AvailableCopies})

		c.JSON(200, gin.H{"message": "Book Copies incremented successfully", "Details": finalData})
	}
}

func removeBook(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	ISBN := c.Request.FormValue("isbn")
	AdminEmail := c.Request.FormValue("adminEmail")

	u64_Id, err := strconv.ParseUint(ISBN, 10, 32)
	if err != nil {
		panic(err)
	}
	isbn := uint(u64_Id)

	//DATABASE OPERATIONS

	DB := db_connection()

	var usr Users
	DB.First(&usr, "email = ?", AdminEmail)
	if usr.Role != "admin" {
		c.JSON(400, gin.H{"message": "Unknown User"})
		return
	}

	var Book BookInventory

	record := DB.First(&Book, "ISBN = ?", isbn)

	if record.Error != nil {
		// fmt.Println("Error: ", record.Error)
		c.JSON(404, gin.H{"message": "Book Not Found", "ISBN": isbn})
	} else {
		if Book.AvailableCopies > 0 {
			// if Book.AvailableCopies > 0 && Book.TotalCopies != Book.AvailableCopies {
			TC := Book.TotalCopies
			// fmt.Println("TC ", TC)
			AC := Book.AvailableCopies
			// fmt.Println("AC ", AC)

			Issued := TC - AC
			// fmt.Println("ISSUED ", Issued)

			TC = TC - 1
			// fmt.Println("TC ", TC)

			AC = TC - Issued
			// fmt.Println("AC ", AC)

			DB.Model(&Book).Where("ISBN = ?", isbn).Select("TotalCopies", "AvailableCopies").Updates(BookInventory{TotalCopies: TC, AvailableCopies: AC})
			c.JSON(200, gin.H{"message": "Removed 1 Book", "ISBN": isbn})

		} else if Book.AvailableCopies == 0 {
			c.JSON(400, gin.H{"message": "Could not delete the book,Available Books are 0", "ISBN": isbn})
		}
	}
}

func updateBook(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	ISBN := c.Request.FormValue("isbn")
	LibID := c.Request.FormValue("libId")
	Title := c.Request.FormValue("title")
	Authors := c.Request.FormValue("authors")
	Publisher := c.Request.FormValue("publisher")
	Version := c.Request.FormValue("version")
	TotalCopies := c.Request.FormValue("totalCopies")
	AdminEmail := c.Request.FormValue("adminEmail")

	u64_Id, err := strconv.ParseUint(ISBN, 10, 32)
	if err != nil {
		panic(err)
	}
	isbn := uint(u64_Id)

	u64_libId, err := strconv.ParseUint(LibID, 10, 32)
	if err != nil {
		panic(err)
	}
	libId := uint(u64_libId)

	u64_version, err := strconv.ParseUint(Version, 10, 32)
	if err != nil {
		panic(err)
	}
	version := uint(u64_version)

	u64_totalCopies, err := strconv.ParseUint(TotalCopies, 10, 32)
	if err != nil {
		panic(err)
	}
	totalCopies := uint(u64_totalCopies)

	finalData := BookInventory{
		ISBN:            isbn,
		LibID:           libId,
		Title:           Title,
		Authors:         Authors,
		Publisher:       Publisher,
		Version:         version,
		TotalCopies:     totalCopies,
		AvailableCopies: totalCopies,
	}

	//DATABASE OPERATIONS

	DB := db_connection()

	var usr Users
	DB.First(&usr, "email = ?", AdminEmail)
	if usr.Role != "admin" {
		c.JSON(400, gin.H{"message": "Unknown User", "Details": finalData})
		return
	}

	var Book BookInventory

	DB.First(&Book, "ISBN = ?", finalData.ISBN)

	// When book is there
	if Book.ISBN == finalData.ISBN {
		DB.Model(&Book).Where("ISBN = ?", finalData.ISBN).Updates(&finalData)

		c.JSON(201, gin.H{"message": "Book Updated successfully", "Details": finalData})
	} else {
		c.JSON(200, gin.H{"message": "Book Not Found", "Details": finalData})
	}
}

func listIssueRequests(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	AdminEmail := c.Param("Email")

	// Container for Storing all rows in Array
	var Book_requests = []RequestEvents{}

	//DATABASE OPERATIONS
	DB := db_connection()

	var usr Users
	DB.First(&usr, "email = ?", AdminEmail)
	if usr.Role != "admin" {
		c.JSON(400, gin.H{"message": "Unknown User"})
		return
	}

	result := DB.Find(&Book_requests)

	// returns found records count, equals `len(users)`
	println(result.RowsAffected)
	// println(result.Error)

	c.IndentedJSON(http.StatusOK, Book_requests)
}

func approveRequest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	RequestId := c.Request.FormValue("requestId")
	ApproverId := c.Request.FormValue("adminEmail")

	//Converting to uint
	u64_Id, err := strconv.ParseUint(RequestId, 10, 32)
	if err != nil {
		panic(err)
	}
	requestId := uint(u64_Id)

	now := time.Now()
	ApproverDate := now.Format("2006-01-02")

	//DATABASE OPERATIONS
	DB := db_connection()

	var usr Users
	DB.First(&usr, "email = ?", ApproverId)
	if usr.Role != "admin" {
		c.JSON(400, gin.H{"message": "Unknown User"})
		return
	}

	var Book_requests = RequestEvents{}
	DB.First(&Book_requests, "req_id = ?", requestId)

	if Book_requests.ReqID == requestId {
		Book_requests.RequestType = "approved"
		Book_requests.ApprovalDate = ApproverDate
		Book_requests.ApproverID = ApproverId

		result_requestApproved := DB.Model(&Book_requests).Where("req_id = ?", requestId).Updates(&Book_requests)
		if result_requestApproved.Error != nil {
			panic(result_requestApproved.Error)
		} else {
			fmt.Println("Request approved")
		}

		var set_issue IssueRegistery
		set_issue.ISBN = Book_requests.BookID
		set_issue.ReaderID = Book_requests.ReaderID
		set_issue.IssueApproverID = Book_requests.ApproverID
		set_issue.IssueStatus = Book_requests.RequestType
		set_issue.IssueDate = Book_requests.ApprovalDate
		// set_issue.ExpectedReturnDate = "NA"
		// set_issue.ReturnDate = ""
		// set_issue.ReturnApproverID = ""

		create_result := DB.Create(&set_issue)

		if create_result.Error != nil {
			panic(create_result.Error)
		}

		c.IndentedJSON(http.StatusOK, "Request Approved")
	}

}

// ******************
// READER API's
// ******************

func searchBookByTitle(c *gin.Context) {
	Title := c.Param("Title")

	//DATABASE OPERATIONS
	DB := db_connection()
	var Book BookInventory
	search_result := DB.First(&Book, "title = ?", Title)
	if search_result.Error != nil {
		panic(search_result.Error)
	}

	if Book.AvailableCopies > 0 {
		c.IndentedJSON(200, Book)
	} else {
		c.IndentedJSON(400, "Book not Avialable")
	}
}

func searchBookByAuthor(c *gin.Context) {
	Author := c.Param("Author")

	//DATABASE OPERATIONS
	DB := db_connection()
	var Book BookInventory
	search_result := DB.First(&Book, "authors = ?", Author)
	if search_result.Error != nil {
		panic(search_result.Error)
	}

	if Book.AvailableCopies > 0 {
		c.IndentedJSON(200, Book)
	} else {
		c.IndentedJSON(400, "Book not Avialable")
	}
}

func searchBookByPublisher(c *gin.Context) {
	Publisher := c.Param("Publisher")

	//DATABASE OPERATIONS
	DB := db_connection()
	var Book BookInventory
	search_result := DB.First(&Book, "publisher = ?", Publisher)
	if search_result.Error != nil {
		panic(search_result.Error)
	}

	if Book.AvailableCopies > 0 {
		c.IndentedJSON(200, Book)
	} else {
		c.IndentedJSON(400, "Book not Avialable")
	}
}

func raiseRequest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extracting Values from submitted form
	BookId := c.Request.FormValue("bookId")
	ReaderId := c.Request.FormValue("readerEmail")

	//Converting to uint
	u64_Id, err := strconv.ParseUint(BookId, 10, 32)
	if err != nil {
		panic(err)
	}
	bookId := uint(u64_Id)

	//DATABASE OPERATIONS
	DB := db_connection()
	var Book_requested BookInventory
	result_bookFound := DB.First(&Book_requested, "ISBN = ?", bookId)
	if result_bookFound.Error != nil {
		panic(result_bookFound.Error)
	}

	if Book_requested.AvailableCopies > 0 {
		var book_request RequestEvents

		book_request.BookID = bookId
		book_request.ReaderID = ReaderId
		now := time.Now()
		RequestDate := now.Format("2006-01-02")
		book_request.RequestDate = RequestDate

		create_result := DB.Create(&book_request)

		if create_result.Error != nil {
			panic(create_result.Error)
		}

		c.IndentedJSON(201, "Issue Requested")
	} else {
		c.IndentedJSON(400, "No Avialable Copies for the Book Requested, Request Rejected")
	}
}
