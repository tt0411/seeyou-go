package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name         string
		Port         string
		TokenTimeout int
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
	Mail struct {
		Smtp     string
		SmtpPort int
		User     string
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("配置文件读取失败: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("无法获取配置: %v", err)
	}

	initDB()
	initRedisDB()

}
