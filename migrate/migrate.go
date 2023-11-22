package main

import (
	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
