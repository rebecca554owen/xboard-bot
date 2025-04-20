package handler

import (
	"fmt"
	"strings"
	"xboard-bot/xboard"

	tele "gopkg.in/telebot.v4"
)

// Handler 管理命令处理器
type Handler struct {
	bot *tele.Bot      // Telegram 机器人实例
	xb  *xboard.Client // XBoard 客户端
}

// NewHandler 创建新的命令处理器
func NewHandler(bot *tele.Bot, xb *xboard.Client) *Handler {
	return &Handler{bot: bot, xb: xb}
}

// Register 设置命令处理器
func (h *Handler) Register() {
	h.bot.Handle("/start", h.handleStart)
	h.bot.Handle("/help", h.handleHelp)
	h.bot.Handle("/bind", h.handleBind)
	h.bot.Handle("/info", h.handleInfo)
	h.bot.Handle("/sub", h.handleSub)
	h.bot.Handle("/traffic", h.handleTraffic)
	h.bot.Handle("/checkin", h.handleCheckin)
	h.bot.Handle("/about", h.handleAbout)
}

// handleStart 处理 /start 命令
func (h *Handler) handleStart(c tele.Context) error {
	return c.Send("欢迎使用 XBoard Bot！请使用 /bind <key> 绑定您的 XBoard 账户。\n使用 /help 查看更多命令。")
}

// handleHelp 处理 /help 命令
func (h *Handler) handleHelp(c tele.Context) error {
	helpText := `可用命令：
/start - 开始使用
/help - 显示帮助
/bind <key> - 绑定 XBoard 账户
/info - 查看账户信息
/sub - 获取订阅链接
/traffic - 查看流量信息
/checkin - 每日签到
/about - 关于机器人`
	return c.Send(helpText)
}

// handleBind 处理 /bind 命令
func (h *Handler) handleBind(c tele.Context) error {
	args := strings.Fields(c.Text())
	if len(args) != 2 {
		return c.Send("用法: /bind <key>")
	}

	key := args[1]
	userID := c.Sender().ID
	if err := h.xb.BindUser(userID, key); err != nil {
		return c.Send(fmt.Sprintf("绑定失败: %v", err))
	}
	return c.Send("绑定成功！")
}

// handleInfo 处理 /info 命令
func (h *Handler) handleInfo(c tele.Context) error {
	userID := c.Sender().ID
	user, err := h.xb.GetUserInfo(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("获取用户信息失败: %v", err))
	}

	info := fmt.Sprintf("用户信息:\n邮箱: %s\n余额: %.2f\n佣金: %.2f\n计划ID: %d",
		user.Email, user.Balance, user.Commission, user.PlanID)
	return c.Send(info)
}

// handleSub 处理 /sub 命令
func (h *Handler) handleSub(c tele.Context) error {
	userID := c.Sender().ID
	sub, err := h.xb.GetSubscription(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("获取订阅链接失败: %v", err))
	}
	return c.Send(sub)
}

// handleTraffic 处理 /traffic 命令
func (h *Handler) handleTraffic(c tele.Context) error {
	userID := c.Sender().ID
	traffic, err := h.xb.GetTraffic(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("获取流量信息失败: %v", err))
	}

	info := fmt.Sprintf("流量信息:\n上传: %.2f MB\n下载: %.2f MB\n总计: %.2f MB\n剩余: %.2f MB",
		traffic.Upload/1024/1024, traffic.Download/1024/1024, traffic.Total/1024/1024, traffic.Remaining/1024/1024)
	return c.Send(info)
}

// handleCheckin 处理 /checkin 命令
func (h *Handler) handleCheckin(c tele.Context) error {
	userID := c.Sender().ID
	result, err := h.xb.CheckIn(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("签到失败: %v", err))
	}
	return c.Send(result)
}

// handleAbout 处理 /about 命令
func (h *Handler) handleAbout(c tele.Context) error {
	return c.Send("XBoard Bot - 基于 XBoard 的 Telegram 机器人\n版本: 1.0.0\n作者: YourName")
}