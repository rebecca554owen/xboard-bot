package handler

import (
	"net/url"
	"strings"
	"xboard-bot/utils/types"

	tele "gopkg.in/telebot.v4"
)

// handleBind 处理 /bind 命令
func (h *Handler) handleBind(c tele.Context) error {
	// 确保是私聊
	if !c.Message().Private() {
		return c.Send("请私聊我进行绑定操作")
	}

	args := strings.Fields(c.Text())
	if len(args) < 2 {
		return c.Send("参数有误，请携带订阅地址发送，格式: /bind <订阅地址>")
	}
	subscriptionAddress := args[1]

	// 解析订阅地址获取token
	parsedUrl, err := url.Parse(subscriptionAddress)
	if err != nil {
		return c.Send("订阅地址无效")
	}

	// 首先尝试从查询参数获取token
	token := ""
	if parsedUrl.RawQuery != "" {
		query := parsedUrl.Query()
		token = query.Get("token")
	}

	// 如果查询参数中没有token，尝试从路径获取
	if token == "" {
		pathParts := strings.Split(strings.Trim(parsedUrl.Path, "/"), "/")
		if len(pathParts) > 0 {
			token = pathParts[len(pathParts)-1]
		}
	}

	if token == "" {
		return c.Send("订阅地址无效")
	}

	// 查询token对应的用户是否存在
	db := h.userRepo.GetDb()
	var user types.User
	err = db.Table("v2_user").Where("token = ?", token).First(&user).Error
	if err != nil || user.ID == 0 {
		return c.Send("token不存在,请检查订阅地址")
	}

	// 检查用户是否已绑定
	if user.TelegramID != 0 {
		if user.TelegramID == c.Sender().ID {
			return c.Reply("您已经绑定过了")
		}
		return c.Send("该账号已经绑定了其他Telegram账号")
	}

	// 执行绑定操作
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user.TelegramID = c.Sender().ID
	if err := tx.Table("v2_user").Where("token = ?", user.Token).Update("telegram_id", c.Sender().ID).Error; err != nil {
		tx.Rollback()
		return c.Send("绑定失败")
	}

	// 验证更新是否成功
	var updatedUser types.User
	if err := tx.Table("v2_user").Where("token = ?", user.Token).First(&updatedUser).Error; err != nil || updatedUser.TelegramID != c.Sender().ID {
		tx.Rollback()
		return c.Send("绑定失败")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Send("绑定失败")
	}

	// 发送绑定成功消息
	return c.Reply("绑定成功")
}
