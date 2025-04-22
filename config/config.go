package config

import (
	"github.com/spf13/viper"
)

// Config 存储应用程序配置
type Config struct {
	Telegram struct {
		Token string `mapstructure:"token"`
	} `mapstructure:"telegram"`

	MySQL struct {
		Host     string `mapstructure:"db_host"`
		Port     int    `mapstructure:"db_port"`
		User     string `mapstructure:"db_username"`
		Password string `mapstructure:"db_password"`
		Database string `mapstructure:"db_database"`
	} `mapstructure:"mysql"`

	Redis struct {
		Enabled  bool   `mapstructure:"redis_enable"`
		Host     string `mapstructure:"redis_host"`
		Port     int    `mapstructure:"redis_port"`
		Password string `mapstructure:"redis_password"`
		Database int    `mapstructure:"redis_db"`
	} `mapstructure:"redis"`
}

// Load 使用 viper 读取并解析 config.yaml 文件
func Load(path string) (*Config, error) {
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AutomaticEnv() // 允许使用环境变量覆盖配置

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
