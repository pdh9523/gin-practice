package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/post"
	postModel "github.com/pdh9523/gin-practice/internal/domain/post/model"
	"github.com/pdh9523/gin-practice/internal/domain/user"
	userModel "github.com/pdh9523/gin-practice/internal/domain/user/model"
	"github.com/pdh9523/gin-practice/internal/infra/db"
)

func SetupRouter() *gin.Engine {
	// DB 마운트
	db.Init()
	db.DB.AutoMigrate(&userModel.User{}, &postModel.Post{})

	r := gin.Default()
	r.Group("/api/v1/")

	user.MountUserRoutes(r)
	post.MountPostRoutes(r)

	return r
}
