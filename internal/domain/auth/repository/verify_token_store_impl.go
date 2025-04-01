package repository

import (
	"fmt"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"time"
)

type VerifyTokenStoreImpl struct {
	CacheStore cache.GlobalCacheStore
}

func NewVerifyTokenStore(cacheStore cache.GlobalCacheStore) VerifyTokenStore {
	return &VerifyTokenStoreImpl{CacheStore: cacheStore}
}

const verifyTokenExpires = time.Minute * 15

func verifyTokenKey(token string) string {
	return fmt.Sprintf("verify_token:%s", token)
}

func (v *VerifyTokenStoreImpl) Save(token, email string) error {
	return v.CacheStore.Save(verifyTokenKey(token), email, verifyTokenExpires)
}

func (v *VerifyTokenStoreImpl) FindEmailByToken(token string) (string, error) {
	return v.CacheStore.Find(verifyTokenKey(token))
}

func (v *VerifyTokenStoreImpl) Delete(token string) error {
	return v.CacheStore.Delete(verifyTokenKey(token))
}
