package cache_service

import (
	"strconv"
	"time"

	"github.com/suthanth/bookstore_users_api/cache"
	"github.com/suthanth/bookstore_users_api/dto/token_dto"
)

type ICacheService interface {
	SaveTokenDetails(token_dto.TokenDetailsDto, uint64) error
	FetchTokenDetails(string) (uint64, error)
	DeleteTokenDetails(string) error
}

type CacheServiceImpl struct {
	RedisClient cache.IRedisClient
}

func NewCacheService(redisClient cache.IRedisClient) *CacheServiceImpl {
	service := &CacheServiceImpl{
		RedisClient: redisClient,
	}
	return service
}

func (c CacheServiceImpl) SaveTokenDetails(tokenDetails token_dto.TokenDetailsDto, userId uint64) error {
	at := time.Unix(tokenDetails.AtExpires, 0)
	rt := time.Unix(tokenDetails.RtExpires, 0)
	now := time.Now()
	_, err := c.RedisClient.Set(tokenDetails.AccessUUID, userId, at.Sub(now))
	if err != nil {
		return err
	}
	_, err = c.RedisClient.Set(tokenDetails.RefreshUUID, userId, rt.Sub(now))
	return err
}

func (c CacheServiceImpl) FetchTokenDetails(accessUUID string) (uint64, error) {
	val, err := c.RedisClient.Get(accessUUID)
	var userId uint64
	if err != nil {
		return userId, err
	}
	userId, err = strconv.ParseUint(val, 10, 64)
	return userId, err
}

func (c CacheServiceImpl) DeleteTokenDetails(accessUUID string) error {
	_, err := c.RedisClient.Del(accessUUID)
	return err
}
