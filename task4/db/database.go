package db

import (
	"blog_system/config"
	"blog_system/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局共享的数据库连接实例

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.Database.User,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	if config.GlobalConfig.Database.Debug {
		DB = DB.Debug()
	}
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
