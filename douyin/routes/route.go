package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 API 路由
	ApiRoutes(router)

	// 配置 404 请求
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理404请求
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error_code":    404,
			"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
		})
	})
}
