package db

import (
	"blog_system/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局共享的数据库连接实例

func InitDB() {
	// todo 连接信息配置化
	dsn := "root:root@tcp(192.168.200.130:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	// 开启Debug日志 todo 配置化
	DB = DB.Debug()
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
