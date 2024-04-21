package main

import (
	"fmt"
	"go.uber.org/zap"
	"serenity-book-api/pkg/global"
	"serenity-book-api/pkg/http"
	"serenity-book-api/pkg/initialization"
	"serenity-book-api/pkg/log"
)

func main() {
	conf := initialization.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+fmt.Sprintf("%d", global.GConfig.System.Addr)))
	logger.Info("docs addr", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger/index.html", global.GConfig.System.Host, global.GConfig.System.Addr)))
	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", global.GConfig.System.Addr))
}
