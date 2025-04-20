package utils

import (
	"log"
	"xboard-bot/xboard"

	tele "gopkg.in/telebot.v4"
)

// Logging 日志记录中间件
func Logging() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			log.Printf("收到来自 %d 的消息: %s", c.Sender().ID, c.Text())
			return next(c)
		}
	}
}

// Binding 绑定检查中间件
func Binding(xb *xboard.Client) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if _, err := xb.GetUserInfo(c.Sender().ID); err != nil {
				return c.Send("请先使用 /bind 命令绑定您的 XBoard 账户")
			}
			return next(c)
		}
	}
}