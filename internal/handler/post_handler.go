package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/db"
	"github.com/pdh9523/gin-practice/internal/model"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []model.Post

	db.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPostById(c *gin.Context) {
	var post model.Post

	id := c.Param("id")
	if err := db.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	var post model.Post
	id := c.Param("id")
	if err := db.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&model.Post{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	}
	c.Status(http.StatusNoContent)
}
