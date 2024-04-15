package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(port string, ginEngine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      ginEngine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func RunServer(port string) {
	// 初始化路由
	ginEngine := RouterInit()
	// 初始化HTTP服务器
	srv := initServer(port, ginEngine)
	// 启动服务
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("listen: %s\n", err))
	}
}
