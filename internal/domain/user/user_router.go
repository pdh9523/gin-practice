package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/user/handler"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/domain/user/service"
	"github.com/pdh9523/gin-practice/internal/infra/db"
)

func MountUserRoutes(r *gin.Engine) {

	userRepository := repository.NewGormUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	route := r.Group("/user")
	{
		route.POST("/register", userHandler.RegisterUser)
	}
}
