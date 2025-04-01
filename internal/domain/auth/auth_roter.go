package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/auth/handler"
	"github.com/pdh9523/gin-practice/internal/domain/auth/repository"
	"github.com/pdh9523/gin-practice/internal/domain/auth/service"
	userRepository "github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"github.com/pdh9523/gin-practice/internal/infra/db"
	"github.com/pdh9523/gin-practice/internal/infra/email"
	"github.com/pdh9523/gin-practice/internal/middleware"
	"os"
)

func MountAuthRoutes(r *gin.Engine, c cache.GlobalCacheStore) {

	userRepo := userRepository.NewGormUserRepository(db.DB)
	refreshTokenStore := repository.NewRefreshTokenStore(c)
	verifyTokenStore := repository.NewVerifyTokenStore(c)
	emailSender := email.NewGoEmailSender(
		os.Getenv("EMAIL_FROM"),
		os.Getenv("EMAIL_HOST"),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		587,
	)
	authService := service.NewAuthService(userRepo, refreshTokenStore, verifyTokenStore, emailSender)
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
