package handler

import (
	"xboard-bot/utils/mysql"

	tele "gopkg.in/telebot.v4"
)

type Handler struct {
	bot *tele.Bot
	userRepo mysql.Repo
}

func NewHandler(bot *tele.Bot, userRepo mysql.Repo) *Handler {
	return &Handler{bot: bot, userRepo: userRepo}
}

// RegisterAll 注册所有命令处理器
func (h *Handler) RegisterAll() {
	h.bot.Handle("/start", h.handleStart)
	h.bot.Handle("/bind", h.handleBind)
	h.bot.Handle("/unbind", h.handleUnbind)
	h.bot.Handle("/subinfo", h.handleSubInfo)
	// 可以在这里添加更多命令注册
}
