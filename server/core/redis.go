package core

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/global"
	"time"
)

func Redis() *redis.Client {
	return InitRedis()
}

// InitRedis DB可以换，0-15
func InitRedis() *redis.Client {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       0,
		PoolSize: redisConfig.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong)
	if err != nil {
		global.Log.Error("redis连接错误", err)
		return nil
	}
	return rdb
}
