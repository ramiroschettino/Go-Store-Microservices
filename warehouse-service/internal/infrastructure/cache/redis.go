package infrastructure

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // No password set
		DB:       0,  // Use default DB
	})
}
