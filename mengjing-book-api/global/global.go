package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mengjing-book/config"
)

// 全局变量
var (
	DB           *gorm.DB            // mysql
	Redis        *redis.Client       // redis
	ServerConfig config.ServerConfig // 项目配置
	Vp           *viper.Viper        // viper
	Logger       *zap.Logger         // 日志
)
