package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"serenity-book-api/docs"
	"serenity-book-api/internal/handler"
	"serenity-book-api/internal/middleware"
	"serenity-book-api/pkg/global"
	"serenity-book-api/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
	systemHandler handler.SystemHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	r.Use(
		middleware.CORSMiddleware(),
	)

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := r.Group(global.GConfig.System.RouterPrefix)
	// 健康监测
	PublicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 基础功能路由
	baseApi := PublicGroup.Group("/base")
	// 验证码接口
	baseApi.GET("/captcha", systemHandler.GetCaptcha)
	// login
	sysApi := PublicGroup.Group("/sys")
	sysApi.GET("/login", userHandler.Login)
	return r
}
