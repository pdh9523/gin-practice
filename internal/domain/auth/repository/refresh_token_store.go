package repository

import "time"

type RefreshTokenStore interface {
	FindByID(userID uint) (string, error)
	Save(userID uint, value string, ttl time.Duration) error
	Delete(userID uint) error
}
