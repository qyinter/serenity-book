package core

import (
	"fmt"
	"github.com/spf13/viper"
	"serenity-book/global"
)

// InitConfigByViper 将配置文件解析为global.ServerConfig
func InitConfigByViper(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	if err = v.Unmarshal(&global.ServerConfig); err != nil {
		fmt.Println(err)
	}

	return v
}
