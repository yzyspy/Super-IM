package core

import (
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis(host string, port int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "", // No password set
		DB:       0,  // Use default DB
	})
	return rdb
}
