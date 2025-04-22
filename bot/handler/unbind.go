package handler

import (
	"xboard-bot/utils/types"

	tele "gopkg.in/telebot.v4"
)

// handleUnbind 处理 /unbind 命令
func (h *Handler) handleUnbind(c tele.Context) error {
	// 确保是私聊
	if !c.Message().Private() {
		return c.Send("请私聊我进行解绑操作")
	}

	// 查询当前用户
	db := h.userRepo.GetDb()
	var user types.User
	err := db.Table("v2_user").Where("telegram_id = ?", c.Sender().ID).First(&user).Error
	if err != nil || user.ID == 0 {
		return c.Send("没有查询到您的用户信息，请先绑定账号")
	}

	// 检查操作权限：仅当前用户可以解绑自己
	if user.TelegramID != c.Sender().ID {
		return c.Send("您没有权限执行此操作")
	}

	// 执行解绑操作
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user.TelegramID = 0
	if err := tx.Table("v2_user").Where("id = ?", user.ID).Update("telegram_id", 0).Error; err != nil {
		tx.Rollback()
		return c.Send("解绑失败")
	}

	// 验证更新是否成功
	var updatedUser types.User
	if err := tx.Table("v2_user").Where("id = ?", user.ID).First(&updatedUser).Error; err != nil || updatedUser.TelegramID != 0 {
		tx.Rollback()
		return c.Send("解绑失败")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Send("解绑失败")
	}

	// 发送解绑成功消息
	return c.Reply("解绑成功")
}