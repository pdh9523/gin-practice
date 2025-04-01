package repository

import (
	"fmt"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"time"
)

type RefreshTokenStoreImpl struct {
	CacheStore cache.GlobalCacheStore
}

func NewRefreshTokenStore(cacheStore cache.GlobalCacheStore) RefreshTokenStore {
	return &RefreshTokenStoreImpl{CacheStore: cacheStore}
}

func key(userID uint) string {
	return fmt.Sprintf("refresh_token:%d", userID)
}

func (r *RefreshTokenStoreImpl) Save(userID uint, value string, ttl time.Duration) error {
	return r.CacheStore.Save(key(userID), value, ttl)
}

func (r *RefreshTokenStoreImpl) FindByID(userID uint) (string, error) {
	return r.CacheStore.Find(key(userID))
}

func (r *RefreshTokenStoreImpl) Delete(userID uint) error {
	return r.CacheStore.Delete(key(userID))
}
