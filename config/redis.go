package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"seeyou-go/global"
)

func initRedisDB() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis 连接失败, %s", err)
	}

	global.RedisDB = RedisClient
}
