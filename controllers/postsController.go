package controllers

import (
	"github.com/asliabhi12/api-task/initializers"
	"github.com/asliabhi12/api-task/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreatePost(c *gin.Context) {
	// get data off req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return

	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	// get id off the url
	id := c.Param("id")

	// get the posts

	var post models.Post
	initializers.DB.Find(&post, id)

	// resoponding with them

	c.JSON(200, gin.H{
		"posts": post,
	})

}



func PostsUpdate (c *gin.Context) {
	// get id off the url
	id := c.Param("id")

// get the data off the req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post we are updating 
	var post models.Post
	initializers.DB.Find(&post, id)

	// update it 
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with it

	c.JSON(200, gin.H{
		"posts": post,
	})
}


