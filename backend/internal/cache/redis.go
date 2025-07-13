package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 密码
		DB:       0,                // 默认DB
	})

	// 测试连接
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("Redis连接失败: %v", err)
		panic(err)
	}

	logrus.Info("Redis连接成功")
}

// AddToBlacklist 将token添加到黑名单
func AddToBlacklist(token string, expiration time.Duration) error {
	err := RedisClient.Set(ctx, "blacklist:"+token, true, expiration).Err()
	if err != nil {
		logrus.Errorf("添加token到黑名单失败: %v", err)
		return err
	}
	return nil
}

// IsInBlacklist 检查token是否在黑名单中
func IsInBlacklist(token string) bool {
	exists, err := RedisClient.Exists(ctx, "blacklist:"+token).Result()
	if err != nil {
		logrus.Errorf("检查token黑名单状态失败: %v", err)
		return false
	}
	return exists > 0
}
