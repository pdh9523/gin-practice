package post

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/post/handler"
)

func MountPostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("", handler.GetPosts)
		post.GET("/:id", handler.GetPostById)
		post.POST("", handler.CreatePost)
		post.DELETE("/:id", handler.DeletePost)
		post.PUT("/:id", handler.UpdatePost)
	}
}
