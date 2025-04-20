package bot

import (
	"log"
	"time"
	"xboard-bot/config"
	"xboard-bot/utils"
	"xboard-bot/xboard"

	tele "gopkg.in/telebot.v4"

)

// Bot 封装 Telegram 机器人及其依赖
type Bot struct {
	teleBot *tele.Bot      // Telegram 机器人实例
	xb      *xboard.Client // XBoard 客户端
}

// NewBot 创建新的机器人实例
func NewBot(cfg *config.Config) (*Bot, error) {
	// 初始化 XBoard 客户端
	xb := xboard.NewClient(cfg.XBoard.Address, cfg.XBoard.Token)

	// 初始化 Telegram 机器人
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
		xb:      xb,
	}

	// 设置命令处理器
	h := utils.NewHandler(teleBot, xb)
	h.Register()

	// 设置中间件
	b.teleBot.Use(utils.Logging())
	b.teleBot.Use(utils.Binding(xb))

	return b, nil
}

// Start 启动机器人
func (b *Bot) Start() {
	log.Printf("机器人 @%s 已启动", b.teleBot.Me.Username)
	b.teleBot.Start()
}