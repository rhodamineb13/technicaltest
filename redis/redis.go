package redisService

import (
	"fmt"
	"technicaltest/config"

	"github.com/go-redis/redis/v8"
)

var RC *redis.Client

func NewRedisClient() {
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(":%s", config.Conf.RedisPort),
		Password: "",
		DB:       0,
	})

	RC = r
}
