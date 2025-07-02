package main

import (
	"blog_system/db"
	"blog_system/middleWare"
	"blog_system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化全局数据库连接
	db.InitDB()
	defer db.CloseDB()

	r := gin.Default()
	r.Use(middleWare.AuthMiddleware())
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
