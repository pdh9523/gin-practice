package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var accessSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
var accessTokenExpireTime = 1 * time.Hour // 1시간

var refreshSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
var refreshTokenExpireTime = 15 * 24 * time.Hour // 15일

func GenerateToken(userID uint, tokenSecret []byte, expireTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tokenSecret)
}

func GenerateAccessToken(userID uint) (string, error) {
	return GenerateToken(userID, accessSecret, accessTokenExpireTime)
}

func GenerateRefreshToken(userID uint) (string, error) {
	return GenerateToken(userID, refreshSecret, refreshTokenExpireTime)
}

func ParseToken(tokenString string, tokenSecret []byte) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	uid := uint(claims["user_id"].(float64))
	return uid, nil
}

func ParseRefreshToken(tokenString string) (uint, error) {
	return ParseToken(tokenString, refreshSecret)
}

func ParseAccessToken(tokenString string) (uint, error) {
	return ParseToken(tokenString, accessSecret)
}
