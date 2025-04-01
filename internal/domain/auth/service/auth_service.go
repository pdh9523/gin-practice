package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/auth/dto"
)

type AuthService interface {
	Login(loginRequestDto dto.LoginRequestDto) (*dto.TokenResponseDto, error)
	Logout(userID uint) error
	TokenRefresh(userID uint, refreshToken string) (*dto.TokenResponseDto, error)
	SendEmail(email string) error
	VerifyEmail(token string) (string, error)
}
