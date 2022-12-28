package cache

import (
	"context"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"time"
)

type UserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error)
}
type authUserCache struct {
	store      UserStore // mysql
	cacheStore Cache     // redis
}

func NewAuthUserCache(store UserStore, cacheStore Cache) *authUserCache {
	return &authUserCache{store: store, cacheStore: cacheStore}
}

func (c *authUserCache) FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error) {
	var user *usermodel.User

	var key string
	userEmail := condition["email"].(string)

	key = "user:" + userEmail // key store in redis

	c.cacheStore.Get(ctx, key, &user)

	// Try to find data of user in cache
	//if err := c.cacheStore.Get(ctx, key, &user); err != nil {
	//	panic(common.NewCustomError(err, "Error when cache.Get() data"))
	//}

	// if data is found in cache, then return the data
	if user != nil {
		return user, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.FindDataWithCondition(ctx, condition)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*2); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err

}
