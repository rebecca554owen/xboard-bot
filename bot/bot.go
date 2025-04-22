package bot

import (
	"log"
	"time"

	"xboard-bot/bot/handler"
	"xboard-bot/bot/middleware"
	"xboard-bot/config"
	"xboard-bot/utils/mysql"
	"xboard-bot/utils/redis"
	tele "gopkg.in/telebot.v4"
)

// Bot 结构体，包含机器人实例
type Bot struct {
	teleBot *tele.Bot
	dbRepo  mysql.Repo
}

// NewBot 创建新的机器人实例
func NewBot(cfg *config.Config) (*Bot, error) {
	// 初始化Mysql数据库
	dbRepo := mysql.GetDbClient(cfg)

	// 初始化Redis
	if cfg.Redis.Enabled {
		redis.GetRedisClient(cfg)
	}

	// 初始化机器人
	pref := tele.Settings{
		Token:  cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	teleBot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	// 创建机器人实例
	b := &Bot{
		teleBot: teleBot,
		dbRepo:  dbRepo,
	}

	// 中间件记录log
	b.teleBot.Use(middleware.Logging())

	// 创建处理器并注册命令
	handler := handler.NewHandler(teleBot, dbRepo)
	handler.RegisterAll()

	return b, nil
}

// Start 启动机器人
func (b *Bot) Start() {
	log.Printf("机器人已启动，用户名 @%s ", b.teleBot.Me.Username)
	b.teleBot.Start()
}

// Stop 停止机器人
func (b *Bot) Stop() error {
	if b.teleBot != nil {
		log.Println("收到终止信号，正在关闭机器人...")
		b.teleBot.Stop()
		log.Println("机器人已成功关闭。")
		return nil
	}
	return nil
}