package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JwtConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Debug    bool   `mapstructure:"debug"`
}

type JwtConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

var GlobalConfig Config

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v", err))
	}

	// 配置热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("检测到配置文件变更:", e.Name)
		reloadConfig()
	})

	// 解析到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Sprintf("解析配置文件失败: %v", err))
	}

	fmt.Println("配置文件加载成功")
}

func reloadConfig() {
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Printf("热更新配置失败: %v\n", err)
	} else {
		fmt.Println("配置热更新成功")
	}
}
