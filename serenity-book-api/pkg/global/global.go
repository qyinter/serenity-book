package global

import (
	"github.com/spf13/viper"
	"serenity-book-api/pkg/config"
	"serenity-book-api/pkg/log"
	"sync"
)

var (
	// GConfig 全局配置
	GConfig config.Server
	// GLog 全局日志
	GLog log.Logger
	// GViper 配置
	GVP *viper.Viper
	//BlackCache 本地缓存
	lock     sync.RWMutex
	DbConfig config.DbConfig
)
