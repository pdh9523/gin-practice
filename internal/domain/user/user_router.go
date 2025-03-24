package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/user/handler"
)

func MountUserRoutes(r *gin.Engine, handler *handler.UserHandler) {
	r.Group("/user")
	{
		r.POST("/login", handler.LoginUser)
		r.POST("/register", handler.RegisterUser)
	}
}
