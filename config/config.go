package config

import (
	"github.com/spf13/viper"
)

// Config 存储应用程序配置
type Config struct {
	Telegram struct {
		Token string `mapstructure:"token"` // Telegram 机器人令牌
	} `mapstructure:"telegram"`
	XBoard struct {
		Address string `mapstructure:"address"` // XBoard API 地址
		Token   string `mapstructure:"token"`   // XBoard API 令牌
	} `mapstructure:"xboard"`
}

// Load 使用 viper 读取并解析 config.yaml 文件
func Load(path string) (*Config, error) {
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// 解析配置到结构体
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}