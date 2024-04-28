package cache

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/config"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	fmt.Println(config.ConfigInfo.Redis.Password)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.ConfigInfo.Redis.Addr,
		Password: config.ConfigInfo.Redis.Password,
		DB:       constants.UserRedisDB,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		klog.Fatalf("redis init error: %v", err)
	}
}
