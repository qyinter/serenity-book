package utils

import (
	"context"
	"go.uber.org/zap"
	"serenity-book/global"
	"time"
)

const ZERO time.Duration = 0

// RedisLPush List中添加元素
func RedisLPush(key string, data interface{}) {
	err := global.Redis.LPush(context.Background(), key, data).Err()
	if err != nil {
		global.Logger.Error("LPush redis fail!", zap.Error(err))
	}
}

// RedisStringSet 添加string元素
func RedisStringSet(key string, data interface{}) {
	err := global.Redis.Set(context.Background(), key, data, ZERO).Err()
	if err != nil {
		global.Logger.Error("redis string set fail!", zap.Error(err))
	}
}

// RedisStringDurationSet 添加string元素并设置过期时间
func RedisStringDurationSet(key string, data interface{}, d time.Duration) {
	err := global.Redis.Set(context.Background(), key, data, d).Err()
	if err != nil {
		global.Logger.Error("redis string set fail!", zap.Error(err))
	}
}

// RedisGet 从redis中获取元素
func RedisGet(key string) string {
	val, err := global.Redis.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("get redis key fail!", zap.Error(err))
		return ""
	}
	return val
}

// RedisKeyIsExist key在redis中是否存在
// 返回true redis中的key存在
func RedisKeyIsExist(key string) bool {
	exist, err := global.Redis.Exists(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("obtain redis key fail!", zap.Error(err))
	}
	global.Logger.Info("obtain redis val", zap.Int("val =", int(exist)))
	return exist == 1
}

// RedisHSet 设置Hset
func RedisHSet(key string, name string, data interface{}) {
	err := global.Redis.HSet(context.Background(), key, name, data).Err()
	if err != nil {
		global.Logger.Error("redis string set fail!", zap.Error(err))
	}
}

// RedisHGet 获取
func RedisHGet(key string, name string) string {
	result, err := global.Redis.HGet(context.Background(), key, name).Result()
	if err != nil {
		global.Logger.Error("redis string set fail!", zap.Error(err))
	}
	return result
}

// RedisKeyIsHExist 检测hash键是否存在
func RedisKeyIsHExist(key string, name string) bool {
	result, err := global.Redis.HExists(context.Background(), key, name).Result()
	if err != nil {
		global.Logger.Error("obtain redis HKey fail!", zap.Error(err))
	}
	return result
}
