package post

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/post/handler"
	"github.com/pdh9523/gin-practice/internal/domain/post/repository"
	"github.com/pdh9523/gin-practice/internal/domain/post/service"
	"github.com/pdh9523/gin-practice/internal/infra/db"
	"github.com/pdh9523/gin-practice/internal/middleware"
)

func MountPostRoutes(r *gin.Engine) {

	postRepository := repository.NewGormPostRepository(db.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	post := r.Group("/posts")
	{
		post.GET("", postHandler.GetPosts)
		post.GET("/:id", postHandler.GetPostByID)

	}

	postWithAuth := r.Group("/posts")
	postWithAuth.Use(middleware.AuthMiddleware())
	{
		postWithAuth.POST("", postHandler.CreatePost)
		postWithAuth.PATCH("/:id", postHandler.UpdatePost)
		postWithAuth.DELETE("/:id", postHandler.DeletePost)
	}
}
