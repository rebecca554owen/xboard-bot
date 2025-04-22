package handler

import (
	"fmt"
	"xboard-bot/utils/types"
	
	tele "gopkg.in/telebot.v4"
)

// handleSubInfo 处理 /subinfo 命令
func (h *Handler) handleSubInfo(c tele.Context) error {
	// 查询当前用户
	db := h.userRepo.GetDb()
	var user types.User
	err := db.Table("v2_user").Where("telegram_id = ?", c.Sender().ID).First(&user).Error
	if err != nil || user.ID == 0 {
		return c.Send("没有查询到您的用户信息，请先发送 `/bind <订阅地址>` - 绑定 XBoard 账户")
	}

	// 格式化流量数据
	transferEnable := formatTraffic(int64(user.TransferEnable))
	up := formatTraffic(int64(user.U))
	down := formatTraffic(int64(user.D))
	remaining := formatTraffic(int64(user.TransferEnable - (user.U + user.D)))

	// 返回流量信息
	text := "🚥流量查询\n———————————————\n套餐流量：" + transferEnable + 
		"\n已用上行：" + up + 
		"\n已用下行：" + down + 
		"\n剩余流量：" + remaining

	return c.Send(text)
}

// formatTraffic 格式化流量数据
func formatTraffic(bytes int64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%dB", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(bytes)/1024)
	} else if bytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(bytes)/(1024*1024))
	}
	return fmt.Sprintf("%.2fGB", float64(bytes)/(1024*1024*1024))
}