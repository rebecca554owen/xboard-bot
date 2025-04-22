package handler

import tele "gopkg.in/telebot.v4"

// handleStart 处理 /start 命令
func (h *Handler) handleStart(c tele.Context) error {
	text := `欢迎使用 XBoard Bot！
版本: 1.0.0
作者: 周宇航

可用命令：
/bind <订阅地址> - 绑定 XBoard 账户`
	return c.Send(text)
}