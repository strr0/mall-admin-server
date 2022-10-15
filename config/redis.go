package config

import (
	"github.com/go-redis/redis"
	"log"
)

var redisdb *redis.Client

type RedisConfig struct {
	Addr     string
	Password string
}

func InitRedis() {
	var redisConfig RedisConfig
	err := settings.UnmarshalKey("redis", &redisConfig)
	if err != nil {
		log.Fatalf("get config failed: %v", err)
	}
	redisdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		log.Fatalf("connect redis failed: %v", err)
	}
}

func CloseRedis() {
	_ = redisdb.Close()
}

func GetRedis() *redis.Client {
	return redisdb
}
