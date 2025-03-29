package service

import (
	"crypto/subtle"
	"errors"
	"github.com/pdh9523/gin-practice/internal/domain/auth/dto"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"github.com/pdh9523/gin-practice/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepo   repository.UserRepository
	TokenStore cache.RefreshTokenStore
}

func NewAuthService(userRepo repository.UserRepository, tokenStore cache.RefreshTokenStore) AuthService {
	return &AuthServiceImpl{
		UserRepo:   userRepo,
		TokenStore: tokenStore,
	}
}

func (s *AuthServiceImpl) Login(loginRequestDto dto.LoginRequestDto) (*dto.TokenResponseDto, error) {
	user, err := s.UserRepo.FindByEmail(loginRequestDto.Email)
	if err != nil {
		// 회원을 찾을 수 없는 경우
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestDto.Password)); err != nil {
		// 비밀번호가 틀린 경우
		return nil, err
	}

	accessToken, _ := jwt.GenerateAccessToken(user.ID)
	refreshToken, _ := jwt.GenerateRefreshToken(user.ID)
	_ = s.TokenStore.Save(user.ID, refreshToken, jwt.RefreshTokenExpireTime)

	return &dto.TokenResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) Logout(userID uint) error {
	return s.TokenStore.Delete(userID)
}

func (s *AuthServiceImpl) TokenRefresh(userID uint, refreshToken string) (*dto.TokenResponseDto, error) {
	cachedToken, err := s.TokenStore.Find(userID)

	if err != nil {
		// 토큰을 못찾은 경우
		return nil, errors.New("token not found")
	}

	if subtle.ConstantTimeCompare([]byte(cachedToken), []byte(refreshToken)) != 1 {
		// 저장된 토큰과 일치하지 않은 경우 (재 로그인으로 인한 토큰 만료?)
		return nil, errors.New("token expired")
	}

	newAccessToken, _ := jwt.GenerateAccessToken(userID)
	newRefreshToken, _ := jwt.GenerateRefreshToken(userID)
	_ = s.TokenStore.Save(userID, newRefreshToken, jwt.RefreshTokenExpireTime)

	return dto.NewTokenResponseDto(newAccessToken, newRefreshToken), nil
}
