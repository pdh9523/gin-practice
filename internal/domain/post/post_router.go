package post

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/post/handler"
)

func MountPostRoutes(r *gin.Engine, handler *handler.PostHandler) {
	post := r.Group("/posts")
	{
		post.GET("", handler.GetPosts)
		post.GET("/:id", handler.GetPostByID)
		post.POST("", handler.CreatePost)
		post.DELETE("/:id", handler.DeletePost)
		post.PATCH("/:id", handler.UpdatePost)
	}
}
