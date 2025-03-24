package dto

import "github.com/pdh9523/gin-practice/internal/domain/user/model"

type UserRequestDto struct {
	Nickname string `json:"nickname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func ToUser(dto UserRequestDto) *model.User {
	return &model.User{
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
