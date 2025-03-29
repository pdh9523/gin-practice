package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/auth/dto"
	"github.com/pdh9523/gin-practice/internal/domain/auth/service"
	"github.com/pdh9523/gin-practice/internal/util"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequestDto dto.LoginRequestDto
	if err := c.ShouldBindJSON(&loginRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponseDto, err := h.Service.Login(loginRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponseDto)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	claims, ok := util.GetAuthClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err := h.Service.Logout(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *AuthHandler) TokenRefresh(c *gin.Context) {
	claims, ok := util.GetAuthClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var refreshRequestDto dto.RefreshRequestDto
	if err := c.ShouldBindJSON(&refreshRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenResponseDto, err := h.Service.TokenRefresh(claims.UserID, refreshRequestDto.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokenResponseDto)
}
