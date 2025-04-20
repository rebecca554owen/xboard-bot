package main

import (
	"log"
	"xboard-bot/bot"
	"xboard-bot/config"
)

// 主函数，程序入口
func main() {
	// 加载配置文件
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化并启动机器人
	b, err := bot.NewBot(cfg)
	if err != nil {
		log.Fatalf("初始化机器人失败: %v", err)
	}

	log.Println("机器人正在运行...")
	b.Start()
}