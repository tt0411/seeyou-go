package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	RedisDB *redis.Client
)