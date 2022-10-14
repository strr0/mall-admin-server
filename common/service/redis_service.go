package service

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisService struct {
	RedisDB *redis.Client
}

func (iService RedisService) Get(key string) string {
	result, err := iService.RedisDB.Get(key).Result()
	if err != nil {
		log.Printf("redis get failed: %v", err)
		return ""
	}
	return result
}

func (iService RedisService) Set(key, value string) error {
	_, err := iService.RedisDB.Set(key, value, time.Second*60).Result()
	return err
}

func (iService RedisService) Expire(key string, time time.Duration) error {
	_, err := iService.RedisDB.Expire(key, time).Result()
	return err
}

func (iService RedisService) Remove(key string) error {
	_, err := iService.RedisDB.Del(key).Result()
	return err
}
