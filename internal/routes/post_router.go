package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/handler"
)

func MountPostRouter(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("", handler.GetPosts)
		post.GET("/:id", handler.GetPostById)
		post.POST("", handler.CreatePost)
		post.DELETE("/:id", handler.DeletePost)
		post.PUT("/:id", handler.UpdatePost)
	}
}
