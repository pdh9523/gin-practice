package dto

import "github.com/pdh9523/gin-practice/internal/domain/user/model"

type UserResponseDto struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func NewUserResponseDto(user *model.User) *UserResponseDto {
	return &UserResponseDto{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
}
