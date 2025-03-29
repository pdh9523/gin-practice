package util

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/pkg/jwt"
)

func GetAuthClaims(c *gin.Context) (jwt.AuthClaims, bool) {
	val, ok := c.Get("auth")
	if !ok {
		return jwt.AuthClaims{}, false
	}
	claims, ok := val.(jwt.AuthClaims)
	return claims, ok
}
