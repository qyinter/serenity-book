package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"mengjing-book/global"
)

func InitRedis() {
	redisCfg := global.ServerConfig.RedisInfo
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.Database, // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("redis连接失败", err)
	} else {
		global.Redis = client
	}
}
