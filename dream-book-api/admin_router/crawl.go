package admin_router

import (
	"dream-book/service"
	"github.com/gin-gonic/gin"
)

func CrawlRoutes(router *gin.RouterGroup) {
	// CrawlSaveService 保存书源 info 信息
	router.POST("/book_source", service.CrawlSaveService)
	// CrawlUpdateService 更新书源 info 信息
	router.POST("/book_source/:id", service.CrawlUpdateService)
	// CrawlDeleteService 获取 info 信息
	router.GET("/book_source/:id", service.CrawlGetService)
	// 分页获取书源 info 信息
	router.POST("/book_source/list", service.CrawlListService)
	// 删除书源 info 信息
	router.DELETE("/book_source/:id", service.CrawlDeleteService)
}
