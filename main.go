package main

import (
	"os"
	"os/signal"
	"syscall"
	"xboard-bot/bot"
	"xboard-bot/config"
)

// 主函数
func main() {
	// 设置信号处理
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// 加载配置文件
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	// 初始化机器人
	b, err := bot.NewBot(cfg)
	if err != nil {
		panic(err)
	}

	// 监听终止信号
	go func() {
		<-stop
		if b != nil {
			b.Stop()
		}
		os.Exit(0)
	}()

	// 启动机器人
	b.Start() 
}