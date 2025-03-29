package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var accessSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
var AccessTokenExpireTime = 1 * time.Hour // 1시간

var refreshSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
var RefreshTokenExpireTime = 15 * 24 * time.Hour // 15일

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, tokenSecret []byte, expireTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tokenSecret)
}

func GenerateAccessToken(userID uint) (string, error) {
	return GenerateToken(userID, accessSecret, AccessTokenExpireTime)
}

func GenerateRefreshToken(userID uint) (string, error) {
	return GenerateToken(userID, refreshSecret, RefreshTokenExpireTime)
}

func ParseToken(tokenString string, tokenSecret []byte) (AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})

	if err != nil {
		return AuthClaims{}, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return AuthClaims{
			UserID: claims.UserID,
			Email:  claims.Email,
			Role:   claims.Role,
		}, nil
	}
	return AuthClaims{}, errors.New("invalid token")
}

func ParseRefreshToken(tokenString string) (AuthClaims, error) {
	return ParseToken(tokenString, refreshSecret)
}

func ParseAccessToken(tokenString string) (AuthClaims, error) {
	return ParseToken(tokenString, accessSecret)
}
