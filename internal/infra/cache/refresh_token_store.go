package cache

import "time"

type RefreshTokenStore interface {
	Save(userID uint, token string, ttl time.Duration) error
	Find(userID uint) (string, error)
	Delete(userID uint) error
}
