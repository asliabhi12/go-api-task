package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func db_connection() *gorm.DB {
	// Connection String
	const dbURL string = "postgres://postgres:postgres@localhost:5432/test3"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&Library{},
		&Users{},
		&BookInventory{},
		&RequestEvents{},
		&IssueRegistery{})

	return db
}
