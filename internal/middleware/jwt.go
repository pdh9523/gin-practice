package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "토큰이 없거나 유효하지 않습니다."})
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		userId, err := jwt.ParseAccessToken(tokenStr)
		//TODO: 엑세스 토큰이 안맞으면 리프레시 토큰부터 받아보기
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다."})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
