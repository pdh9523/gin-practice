package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/user/dto"
	"github.com/pdh9523/gin-practice/internal/domain/user/service"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userRequestDto dto.UserRequestDto
	if err := c.ShouldBindJSON(&userRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := h.Service.RegisterUser(userRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": dto.NewUserResponseDto(newUser)})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var userLoginDto dto.UserLoginDto
	if err := c.ShouldBindJSON(&userLoginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.Service.LoginUser(userLoginDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, dto.NewUserResponseDto(user))

}
