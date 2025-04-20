package handler

import (
	"log"
	"strings"
	"xboard-bot/xboard"

	tele "gopkg.in/telebot.v4"
)

// LoggingMiddleware 记录传入的更新
func LoggingMiddleware() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			log.Printf("收到来自 %d 的更新: %s", c.Sender().ID, c.Text())
			return next(c)
		}
	}
}

// BindingMiddleware 检查用户是否绑定了 XBoard 账户
func BindingMiddleware(xb *xboard.Client) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			// 跳过 /start, /help, /bind, /about 命令
			if c.Text() == "/start" || c.Text() == "/help" || strings.HasPrefix(c.Text(), "/bind") || c.Text() == "/about" {
				return next(c)
			}

			userID := c.Sender().ID
			if !xb.IsUserBound(userID) {
				return c.Send("请先使用 /bind <key> 绑定您的 XBoard 账户。")
			}
			return next(c)
		}
	}
}