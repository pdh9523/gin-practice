package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/user/dto"
	"github.com/pdh9523/gin-practice/internal/domain/user/model"
)

type UserService interface {
	RegisterUser(userRequestDto dto.UserRequestDto) (*model.User, error)
}
