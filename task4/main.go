package main

import (
	"blog_system/config"
	"blog_system/db"
	"blog_system/logger"
	"blog_system/middleWare"
	"blog_system/routes"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化全局数据库连接
	db.InitDB()
	defer db.CloseDB()
	// 初始化日志工具
	logger.InitLogger()
	r := gin.Default()
	r.Use(middleWare.AuthMiddleware())
	routes.RegisterRoutes(r)
	r.Run(":" + strconv.Itoa(config.GlobalConfig.App.Port))
}
