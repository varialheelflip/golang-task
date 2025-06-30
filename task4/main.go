package main

import (
	"blog_system/db"
)

func main() {
	// 初始化全局数据库连接
	db.InitDB()
	defer db.CloseDB()
}
