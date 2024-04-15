package main

import (
	"go.uber.org/zap"
	"mengjing-book/core"
	"mengjing-book/global"
)

func main() {
	// 读取yaml配置文件中的数据信息
	global.Vp = core.InitConfigByViper("./config/config-dev.yaml")
	// 初始化zapLogger
	global.Logger = core.InitZapLogger()
	zap.ReplaceGlobals(global.Logger)
	// 初始化数据源
	global.DB = core.InitGormMysql()
	// 初始化redis
	core.InitRedis()
	// 数据库连接成功
	if global.DB != nil {
		// core.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	// 运行服务
	core.RunServer(global.ServerConfig.ServiceInfo.Port)
}
