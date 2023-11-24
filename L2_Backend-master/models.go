package main

type Library struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Users struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Email         string `gorm:"not null;uniqueIndex"`
	ContactNumber string
	Role          string
	LibID         uint
	Password      string `gorm:"not null"`
}

type BookInventory struct {
	ISBN            uint `gorm:"primaryKey"`
	LibID           uint
	Title           string
	Authors         string
	Publisher       string
	Version         uint
	TotalCopies     uint
	AvailableCopies uint
}

type RequestEvents struct {
	ReqID        uint `gorm:"primaryKey"`
	BookID       uint
	ReaderID     string
	RequestDate  string
	ApprovalDate string
	ApproverID   string
	RequestType  string
}

type IssueRegistery struct {
	IssueID            uint `gorm:"primaryKey"`
	ISBN               uint
	ReaderID           string
	IssueApproverID    string
	IssueStatus        string
	IssueDate          string
	ExpectedReturnDate string
	ReturnDate         string
	ReturnApproverID   string
}
