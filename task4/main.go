package main

import (
	"blog_system/db"
	"blog_system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化全局数据库连接
	db.InitDB()
	defer db.CloseDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
