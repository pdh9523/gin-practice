package post

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/db"
	"github.com/pdh9523/gin-practice/internal/domain/post/handler"
	"github.com/pdh9523/gin-practice/internal/domain/post/repository"
	"github.com/pdh9523/gin-practice/internal/domain/post/service"
)

func MountPostRoutes(r *gin.Engine) {

	postRepository := repository.NewGormPostRepository(db.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	post := r.Group("/posts")
	{
		post.GET("", postHandler.GetPosts)
		post.GET("/:id", postHandler.GetPostByID)
		post.POST("", postHandler.CreatePost)
		post.DELETE("/:id", postHandler.DeletePost)
		post.PATCH("/:id", postHandler.UpdatePost)
	}
}
