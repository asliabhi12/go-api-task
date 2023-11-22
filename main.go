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

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)


	router.Run()
}
