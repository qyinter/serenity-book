package core

import (
	"dream-book/global"
	"dream-book/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitGormMysql 初始化mysql
func InitGormMysql() *gorm.DB {
	dbInfo := global.ServerConfig.MysqlInfo

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                                             // Slow SQL threshold
			LogLevel:                  utils.GetLoggerLevel(global.ServerConfig.LogInfo.Level), // Log level
			IgnoreRecordNotFoundError: true,                                                    // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                                                    // Don't include params in the SQL log
			Colorful:                  false,                                                   // Disable color
		},
	)

	if dbInfo.DbName == "" {
		log.Fatalln("mysql数据库名不能为空")
		return nil
	}
	// 格式化mysql连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", dbInfo.UserName, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.DbName, dbInfo.Config)

	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		log.Fatalln("mysql数据库连接失败", err)
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+dbInfo.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbInfo.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbInfo.MaxOpenConns)
		return db
	}
}
