package redis

import (
	"bluebell/settings"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	Nil    = redis.Nil
	ctx    = context.Background()
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:           cfg.DB, // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
