package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/user/handler"
)

func MountUserRoutes(r *gin.Engine, handler *handler.UserHandler) {
	route := r.Group("/user")
	{
		route.POST("/login", handler.LoginUser)
		route.POST("/register", handler.RegisterUser)
	}
}
