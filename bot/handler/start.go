package handler

import tele "gopkg.in/telebot.v4"

// handleStart 处理 /start 命令
func (h *Handler) handleStart(c tele.Context) error {
	menu := &tele.ReplyMarkup{}
	btnBind := menu.Data("绑定账户", "bind", "订阅地址")
	btnUnbind := menu.Data("解绑账户", "unbind")
	btnSubinfo := menu.Data("查看订阅信息", "subinfo")
	menu.Inline(
		menu.Row(btnBind),
		menu.Row(btnUnbind),
		menu.Row(btnSubinfo),
	)
	text := `欢迎使用 XBoard Bot！
版本: 1.1.0
作者: 周宇航
请选择：`
	return c.Send(text, menu)
}

func (h *Handler) onBindBtn(c tele.Context) error {
    return c.Edit("请发送 /bind <订阅地址> 格式如下：\n/bind https://xboard.com/s/xxx")
}

func (h *Handler) onUnbindBtn(c tele.Context) error {
    if err := c.Edit("正在处理解绑请求..."); err != nil {
        return err
    }
    return h.handleUnbind(c)
}

func (h *Handler) onSubinfoBtn(c tele.Context) error {
    if err := c.Edit("正在获取流量信息..."); err != nil {
        return err
    }
    return h.handleSubInfo(c)
}