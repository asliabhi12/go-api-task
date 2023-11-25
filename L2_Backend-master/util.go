package main

import (
	"github.com/gin-gonic/gin"
)

func routing() {

	route := gin.Default()

	route.GET("/", Start)
	route.POST("/createLibrary", createLibrary)
	route.POST("/createUser", createUser)
	route.POST("/login", login)
	route.POST("/addBooks", addBooks)
	route.POST("/removeBook", removeBook)
	route.PUT("/updateBook", updateBook)
	route.GET("/listIssueRequests/:Email", listIssueRequests)
	route.POST("/approveRequest", approveRequest)
	route.GET("/searchBookByTitle/:Title", searchBookByTitle)
	route.GET("/searchBookByAuthor/:Author", searchBookByAuthor)
	route.GET("/searchBookByPublisher/:Publisher", searchBookByPublisher)
	route.POST("/raiseRequest", raiseRequest)

	route.Run() // listen and serve on 0.0.0.0:8080
}
