package controllers

import (
	"managing_data/config"
	"managing_data/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.SetupDatabase().Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// retrieves a single post by its ID
func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.SetupDatabase().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func CreatePost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	config.SetupDatabase().Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// updates an existing post by its ID
func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.SetupDatabase().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.SetupDatabase().Model(&post).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.SetupDatabase().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	config.SetupDatabase().Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
