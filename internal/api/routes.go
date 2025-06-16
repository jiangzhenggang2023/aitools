package api

import (
	"aitools/config"
	"aitools/internal/controller"

	"github.com/gin-gonic/gin"
)

// 定义“/v1/document/upload”路由，并配置middleware
// SetupRouter initializes the Gin router and sets up the routes.

// 将路由路径都放在config文件中，方便管理和修改
// 这里的handler是一个自定义的包，包含了处理请求的函数
// SetupRouter initializes the Gin router and sets up the routes.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.GET(config.HelloPath, handler.HelloHandler)
	// r.GET(config.HealthPath, handler.HealthHandler)
	// r.POST(config.UploadPath, handler.UploadHandler, handler.AuthMiddleware())
	r.POST(config.UploadPath, controller.UploadHandler)
	return r
}
