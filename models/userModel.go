package models

type Book struct {
	ISBN            int    `gorm:"primaryKey;column:isbn" json:"isbn"`
	LibID           int    `json:"libId"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"totalCopies"`
	AvailableCopies int    `json:"availableCopies"`
}

type User struct {
	ID            int    `gorm:"primaryKey;column:id" json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"UNIQUE;type:text;not null"`
	Password      string `json:"password"`
	ContactNumber string `json:"contactNumber"`
	Role          string `json:"role"`
	LibID         int    `json:"libId"`
}

type RequestEvents struct {
	ReqID        int    `gorm:"primaryKey;column:reqId" json:"reqId"`
	BookID       int    `gorm:"foreignKey:ISBN" json:"bookId"`
	ReaderID     int    `gorm:"foreignKey:ID" json:"readerId"`
	RequestDate  string `json:"requestDate"`
	ApprovalDate string `json:"approvalDate"`
	ApproverID   int    `json:"approverId"`
	RequestType  string `json:"requestType"`
}

type IssueRegistery struct {
	IssueID            int    `gorm:"primaryKey;column:issueId" json:"issueId"`
	ISBN               int    `gorm:"foreignKey:ISBN" json:"isbn"`
	ReaderID           int    `gorm:"foreignKey:ID" json:"readerId"`
	IssueApproverID    int    `gorm:"foreignKey:ID" json:"issueApproverId"`
	IssueStatus        string `json:"issueStatus"`
	IssueDate          string `json:"issueDate"`
	ExpectedReturnDate string `json:"expectedReturnDate"`
	ReturnDate         string `json:"returnDate"`
	ReturnApproverID   int    `json:"returnApproverId"`
}

type Library struct {
	ID   int    `json:"Id" gorm:"primaryKey;"`
	Name string `json:"name" gorm:"UNIQUE;type:text;not null"`
}
