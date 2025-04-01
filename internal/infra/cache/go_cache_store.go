package cache

import (
	"errors"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type GoCacheTokenStore struct {
	cache *gocache.Cache
}

func NewGoCacheStore(defaultTTL, cleanupInterval time.Duration) GlobalCacheStore {
	return &GoCacheTokenStore{
		cache: gocache.New(defaultTTL, cleanupInterval),
	}
}

func (g *GoCacheTokenStore) Save(key, value string, ttl time.Duration) error {
	g.cache.Set(key, value, ttl)
	return nil
}

func (g *GoCacheTokenStore) Find(key string) (string, error) {
	val, found := g.cache.Get(key)
	if !found {
		return "", errors.New("cache not found")
	}
	return val.(string), nil
}

func (g *GoCacheTokenStore) Delete(key string) error {
	g.cache.Delete(key)
	return nil
}
