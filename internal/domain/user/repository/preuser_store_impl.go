package repository

import (
	"encoding/json"
	"github.com/pdh9523/gin-practice/internal/domain/user/model"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"time"
)

type PreUserStoreImpl struct {
	CacheStore cache.GlobalCacheStore
}

func NewPreUserStore(cacheStore cache.GlobalCacheStore) PreUserStore {
	return &PreUserStoreImpl{
		CacheStore: cacheStore,
	}
}

const (
	PreUserKey     = "pre_user:"
	PreUserExpires = time.Hour * 24 * 15
)

func (c *PreUserStoreImpl) Save(preUser *model.User) error {
	bytes, err := json.Marshal(preUser)
	if err != nil {
		return err
	}
	key := PreUserKey + preUser.Email
	return c.CacheStore.Save(key, string(bytes), PreUserExpires)
}

func (c *PreUserStoreImpl) FindByEmail(email string) (*model.User, error) {
	key := PreUserKey + email
	value, err := c.CacheStore.Find(key)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
