package service

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pdh9523/gin-practice/internal/domain/auth/dto"
	authRepository "github.com/pdh9523/gin-practice/internal/domain/auth/repository"
	userRepository "github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"github.com/pdh9523/gin-practice/internal/infra/email"
	"github.com/pdh9523/gin-practice/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type AuthServiceImpl struct {
	UserRepo          userRepository.UserRepository
	RefreshTokenStore authRepository.RefreshTokenStore
	VerifyTokenStore  authRepository.VerifyTokenStore
	EmailSender       email.EmailSender
}

func NewAuthService(userRepo userRepository.UserRepository, refreshTokenStore authRepository.RefreshTokenStore, verifyTokenStore authRepository.VerifyTokenStore, emailSender email.EmailSender) AuthService {
	return &AuthServiceImpl{
		UserRepo:          userRepo,
		RefreshTokenStore: refreshTokenStore,
		VerifyTokenStore:  verifyTokenStore,
		EmailSender:       emailSender,
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
	_ = s.RefreshTokenStore.Save(user.ID, refreshToken)

	return &dto.TokenResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) Logout(userID uint) error {
	return s.RefreshTokenStore.Delete(userID)
}

func (s *AuthServiceImpl) TokenRefresh(userID uint, refreshToken string) (*dto.TokenResponseDto, error) {
	cachedToken, err := s.RefreshTokenStore.FindByID(userID)

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
	_ = s.RefreshTokenStore.Save(userID, newRefreshToken)

	return dto.NewTokenResponseDto(newAccessToken, newRefreshToken), nil
}

func (s *AuthServiceImpl) SendEmail(email string) error {
	token := uuid.NewString()

	if err := s.VerifyTokenStore.Save(token, email); err != nil {
		return err
	}

	link := fmt.Sprintf("%s/auth/verify?token=%s", os.Getenv("BASE_URL"), token)
	body := fmt.Sprintf("아래 링크를 클릭해 인증을 완료하세요:\n\n👉 %s", link)
	return s.EmailSender.Send(email, "이메일 인증 요청", body)
}

func (s *AuthServiceImpl) VerifyEmail(token string) (string, error) {
	userEmail, err := s.VerifyTokenStore.FindEmailByToken(token)
	if err != nil {
		return "", err
	}
	//TODO: 유저 임시 저장소 만들기
	return userEmail, nil
}
