package cache

import (
	"os"
	"sync"
	"time"

	"github.com/suthanth/bookstore_users_api/logger"

	"github.com/go-redis/redis/v7"
)

var redisClient *RedisClientImp
var once sync.Once

type IRedisClient interface {
	Get(string) (string, error)
	Set(string, interface{}, time.Duration) (string, error)
	Del(keys ...string) (int64, error)
}

type RedisClientImp struct {
	RedisClient *redis.Client
}

func GetRedisClient() *RedisClientImp {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   0,
		})
		pingResponse, err := client.Ping().Result()
		if err != nil {
			logger.SugarLogger.Error("Error setting up redis")
			os.Exit(1)
		}
		logger.SugarLogger.Infof("Pinged redis server. Response %s", pingResponse)
		redisClient = &RedisClientImp{client}
	})
	return redisClient
}

func (r RedisClientImp) Get(key string) (string, error) {
	return r.RedisClient.Get(key).Result()
}

func (r RedisClientImp) Set(key string, val interface{}, ttl time.Duration) (string, error) {
	return r.RedisClient.Set(key, val, ttl).Result()
}

func (r RedisClientImp) Del(key ...string) (int64, error) {
	return r.RedisClient.Del(key...).Result()
}
