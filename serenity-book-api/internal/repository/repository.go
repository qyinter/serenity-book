package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"serenity-book-api/pkg/global"
	"serenity-book-api/pkg/log"
)

type Repository struct {
	db     *gorm.DB
	Rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		db:     db,
		Rdb:    rdb,
		logger: logger,
	}
}

func NewDb() *gorm.DB {
	// TODO: init db
	//db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//return db
	return &gorm.DB{}
}

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         global.DbConfig.Redis.Addr,         // Redis 服务器地址
		Password:     global.DbConfig.Redis.Password,     // Redis 密码
		DB:           global.DbConfig.Redis.DB,           // Redis 数据库
		ReadTimeout:  global.DbConfig.Redis.ReadTimeout,  // 读取超时
		WriteTimeout: global.DbConfig.Redis.WriteTimeout, // 写入超时
	})
	// 测试连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
