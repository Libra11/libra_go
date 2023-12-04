package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"libra.com/user/config"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
}

func (r RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	return r.rdb.Set(ctx, key, value, expire).Err()
}

func (r RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

var Rc *RedisCache

func init() {
	println("init redis")
	config.C = config.InitConfig()
	rdb := redis.NewClient(config.C.InitRedisOptions())
	Rc = &RedisCache{
		rdb: rdb,
	}
}
