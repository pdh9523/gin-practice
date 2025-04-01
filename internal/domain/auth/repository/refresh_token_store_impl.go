package repository

import (
	"fmt"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"github.com/pdh9523/gin-practice/pkg/jwt"
)

type RefreshTokenStoreImpl struct {
	CacheStore cache.GlobalCacheStore
}

func NewRefreshTokenStore(cacheStore cache.GlobalCacheStore) RefreshTokenStore {
	return &RefreshTokenStoreImpl{CacheStore: cacheStore}
}

func refreshTokenKey(userID uint) string {
	return fmt.Sprintf("refresh_token:%d", userID)
}

func (r *RefreshTokenStoreImpl) Save(userID uint, value string) error {
	return r.CacheStore.Save(refreshTokenKey(userID), value, jwt.RefreshTokenExpireTime)
}

func (r *RefreshTokenStoreImpl) FindByID(userID uint) (string, error) {
	return r.CacheStore.Find(refreshTokenKey(userID))
}

func (r *RefreshTokenStoreImpl) Delete(userID uint) error {
	return r.CacheStore.Delete(refreshTokenKey(userID))
}
