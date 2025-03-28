package cache

import (
	"errors"
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type GoCacheTokenStore struct {
	cache *gocache.Cache
}

func NewGoCacheTokenStore(defaultTTL, cleanupInterval time.Duration) RefreshTokenStore {
	return &GoCacheTokenStore{
		cache: gocache.New(defaultTTL, cleanupInterval),
	}
}

func (g *GoCacheTokenStore) key(userID uint) string {
	return "refresh_token:" + fmt.Sprint(userID)
}

func (g *GoCacheTokenStore) Save(userID uint, token string, ttl time.Duration) error {
	key := g.key(userID)
	g.cache.Set(key, token, ttl)
	return nil
}

func (g *GoCacheTokenStore) Find(userID uint) (string, error) {
	key := g.key(userID)
	val, found := g.cache.Get(key)
	if !found {
		return "", errors.New("token not found")
	}
	return val.(string), nil
}

func (g *GoCacheTokenStore) Delete(userID uint) error {
	g.cache.Delete(g.key(userID))
	return nil
}
