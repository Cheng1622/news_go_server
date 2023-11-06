package redis

import (
	"context"
	"os"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

// InitRedis Redis连接池配置
func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Database,
	})

	if str, err := redisClient.Ping(context.Background()).Result(); err != nil || str != "PONG" {
		clog.Log.Panicf("初始化redis数据库异常:", err)
		os.Exit(1)
	}

	clog.Log.Infof("初始化redis数据库完成!")
}