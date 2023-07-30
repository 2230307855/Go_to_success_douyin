package main

import (
	"douyin/config"
	"douyin/dao"
	"douyin/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	// 数据库初始化
	dao.SetupDB()

	// 初始化路由绑定
	routes.SetupRoute(router)

	// 初始化配置
	config.InitConfig()

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
