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

	router.POST("/library", controllers.CreateLibrary) // w
	router.GET("/library", controllers.GetAllLibrary) // w
	router.POST("/book", controllers.CreateBook) // w
	router.GET("/books", controllers.BooksIndex) // w
	router.GET("/book/:id", controllers.BookShow) // w
	router.PUT("/book/:id", controllers.BooksUpdate) // w
	router.POST("/signup", controllers.Signup) // w
	router.POST("/login", controllers.Login) // w
	router.GET("/validate", middleware.RequireAuth, controllers.Validate) // w
	router.POST("/request", controllers.CreateRequest) // w
	router.GET("/requests/", controllers.GetAllRequest) // w
	router.GET("/request/:reqid", controllers.GetRequest) 
	router.PUT("/request/:reqid/", controllers.UpdateRequestByReqID) // doesn't work


	router.Run()
}
# https://github.com/safak/youtube2022/tree/blog-app
