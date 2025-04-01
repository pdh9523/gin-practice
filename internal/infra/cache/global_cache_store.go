package cache

import "time"

type GlobalCacheStore interface {
	Save(key, value string, ttl time.Duration) error
	Find(key string) (string, error)
	Delete(key string) error
}
