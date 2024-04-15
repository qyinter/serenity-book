package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mengjing-book/global"
	"mengjing-book/utils"
	"os"
	"path/filepath"
	"time"
)

const (
	logTmFmtWithMS = "2006/01/02 - 15:04:05"
)

func InitZapLogger() *zap.Logger {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(logTmFmtWithMS))
	}
	path := global.ServerConfig.LogInfo.Path
	filename := global.ServerConfig.LogInfo.Filename

	// 创建日志目录
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	//// 创建日志文件
	filePath := filepath.Join(path, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			panic(fmt.Sprintf("Failed to create log file: %s", err.Error()))
		}
	}

	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(utils.GetLogLevel(global.ServerConfig.LogInfo.Level)),
		OutputPaths:      []string{filePath},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			CallerKey:      "caller", // 打印文件名和行数
			LevelKey:       "level",
			MessageKey:     "msg",
			TimeKey:        "ts",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeTime:     customTimeEncoder, // 自定义时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeName:     zapcore.FullNameEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("zapLogger init fail: %s", err.Error()))
	}
	return logger
}
