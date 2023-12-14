package utils

import (
	"blog-server/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Redis 封装 redis 操作，统一处理错误
var Redis = new(_redis)

type _redis struct{}

// InitRedis 初始化 redis 连接
func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
	// 测试连接状况
	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("Redis 连接失败: ", err)
	}
	log.Println("Redis 连接成功: ", result)

	return rdb
}

// Do 执行指令
func (*_redis) Do(args ...any) (any, error) {
	return rdb.Do(ctx, args...), nil
}

// Keys 根据正则获取keys
func (*_redis) Keys(pattern string) []string {
	return rdb.Keys(ctx, pattern).Val()
}
