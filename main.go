package main

import (
	"github.com/asliabhi12/api-task/controllers"
	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {

	router := gin.Default()

	router.POST("/library", controllers.CreateLibrary) // 
	router.GET("/library", controllers.GetAllLibrary) //

	router.POST("/posts", controllers.CreateBook) 
	router.GET("/posts", controllers.BooksIndex)
	router.GET("/posts/:id", controllers.BookShow)
	router.PUT("/posts/:id", controllers.BooksUpdate)
	router.POST("/signup", controllers.Signup) // w
	router.POST("/login", controllers.Login) // w
	router.GET("/validate", middleware.RequireAuth, controllers.Validate) // w


	router.Run()
}
