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

	router.POST("/library", middleware.OwnerAuth, controllers.CreateLibrary)               // w
	router.GET("/library", middleware.OwnerAuth, controllers.GetAllLibrary)                // w
	router.POST("/book", controllers.CreateBook)                     // w
	router.GET("/books", controllers.BooksIndex)                                           // w
	router.GET("/book/:id", controllers.BookShow)                                          // w
	router.DELETE("/book", middleware.AdminAuth, controllers.RemoveBook)                      // need wrk
	router.PUT("/book/:id", middleware.AdminAuth, controllers.BooksUpdate)                 // w
	router.POST("/signup", controllers.Signup)                                             // w
	router.POST("/login", controllers.Login)                                               // w
	router.POST("/logout", controllers.Logout)
	router.GET("/validate", middleware.OwnerAuth, controllers.Validate)                    // w
	router.POST("/request", controllers.CreateRequest)                                     // w
	router.GET("/requests/", controllers.GetAllRequest)              // w
	router.GET("/request/:reqid", middleware.AdminAuth, controllers.GetRequest)            //w
	router.PUT("/request/:reqid/", middleware.AdminAuth, controllers.UpdateRequestByReqID) // w

	router.POST("/approve-request", middleware.AdminAuth, controllers.ApproveRequest)
	router.GET("/issues", middleware.AdminAuth, controllers.IssuesIndex)


	router.Run()
}
