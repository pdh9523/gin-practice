package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/db"
	"github.com/pdh9523/gin-practice/internal/domain/user/handler"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/domain/user/service"
)

func MountUserRoutes(r *gin.Engine) {

	userRepository := repository.NewGormUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	route := r.Group("/user")
	{
		route.POST("/login", userHandler.LoginUser)
		route.POST("/register", userHandler.RegisterUser)
	}
}
