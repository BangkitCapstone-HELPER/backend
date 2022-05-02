package lib

import (
	"fmt"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cache *redis.Client
}

func NewCache(cfg config.CacheConfig) Cache {
	fmt.Println(cfg.Address())
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password(), // no password set
		DB:       0,              // use default DB
	})
	return Cache{
		Cache: rdb,
	}
}
