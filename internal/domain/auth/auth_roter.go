package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/auth/handler"
	"github.com/pdh9523/gin-practice/internal/domain/auth/service"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"github.com/pdh9523/gin-practice/internal/infra/db"
	"github.com/pdh9523/gin-practice/internal/middleware"
	"time"
)

func MountAuthRoutes(r *gin.Engine) {

	userRepo := repository.NewGormUserRepository(db.DB)
	tokenStore := cache.NewGoCacheTokenStore(0, 10*time.Minute)
	authService := service.NewAuthService(userRepo, tokenStore)
	authHandler := handler.NewAuthHandler(authService)

	routeWithAuth := r.Group("/auth")
	routeWithAuth.Use(middleware.AuthMiddleware())
	{
		routeWithAuth.POST("/logout", authHandler.Logout)
		routeWithAuth.POST("/refresh", authHandler.TokenRefresh)
	}
	route := r.Group("/auth")
	{
		route.POST("/login", authHandler.Login)
	}

}
