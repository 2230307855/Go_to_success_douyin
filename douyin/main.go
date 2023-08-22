package main

import (
	"douyin/dao"
	"douyin/routes"
	"douyin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	// 数据库初始化
	dao.SetupDB()

	// 初始化日志
	utils.InitLogger()
	defer utils.SugarLogger.Sync()

	// 初始化路由绑定
	routes.SetupRoute(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
