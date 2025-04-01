package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/user/dto"
)

type UserService interface {
	RegisterUser(userRequestDto dto.UserRequestDto) (*string, error)
}
