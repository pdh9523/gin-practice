package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/db"
	"github.com/pdh9523/gin-practice/internal/domain/post"
	postModel "github.com/pdh9523/gin-practice/internal/domain/post/model"
	"github.com/pdh9523/gin-practice/internal/domain/user"
	uh "github.com/pdh9523/gin-practice/internal/domain/user/handler"
	userModel "github.com/pdh9523/gin-practice/internal/domain/user/model"
	ur "github.com/pdh9523/gin-practice/internal/domain/user/repository"
	us "github.com/pdh9523/gin-practice/internal/domain/user/service"
)

func SetupRouter() *gin.Engine {
	// DB 마운트
	db.Init()
	db.DB.AutoMigrate(&userModel.User{}, &postModel.Post{})

	r := gin.Default()
	r.Group("/api/v1/")

	userRepository := ur.NewGormUserRepository(db.DB)
	userService := us.NewUserService(userRepository)
	userHandler := uh.NewUserHandler(userService)
	user.MountUserRoutes(r, userHandler)

	post.MountPostRoutes(r)
	return r
}
