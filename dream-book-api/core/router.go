package core

import (
	adminRouter "dream-book/admin_router"
	apiRouter "dream-book/api_router"
	"dream-book/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// RouterInit 整个路由初始化入口
func RouterInit() *gin.Engine {
	//关闭gin日志打印
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = ioutil.Discard
	// 初始化gin服务
	r := gin.Default()
	// 设置模板文件的位置
	r.Use(middleware.CORS())
	// 注册Logger中间件
	r.Use(middleware.RegisterZapLogger())
	//统一异常拦截
	r.Use(middleware.Recover)
	//注册路由

	// 创建API路由组
	api := r.Group("/api")
	apiRouter.RegisterBookRoutes(api)
	// admin 路由组
	admin := r.Group("/admin")
	adminRouter.CrawlRoutes(admin)
	return r
}
