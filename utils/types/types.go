package types

import (
	"time"
)

// 定义 user 结构体
type User struct {
	ID               int    `json:"id"`               		// 用户
	InviteUserID     int    `json:"invite_user_id"`   		// 邀请用户 ID
	TelegramID       int64  `json:"telegram_id"`      		// Telegram ID
	Email            string `json:"email"`            		// 邮箱
	Password         string `json:"password"`         		// 密码
	PasswordAlgo     string `json:"password_algo"`    		// 密码算法
	PasswordSalt     string `json:"password_salt"`    		// 密码盐值
	Balance          int    `json:"balance"`          		// 余额
	Discount         int    `json:"discount"`         		// 折扣
	CommissionType   int    `json:"commission_type"`  		// 佣金类型: 0-system 1-period 2-onetime, tinyint(4), default 0
	CommissionRate   int    `json:"commission_rate"`  		// 佣金率
	CommissionBalance int   `json:"commission_balance"` 	// 佣金余额
	T                int64  `json:"t"`               		// 上次在线时间
	U                int64  `json:"u"`                  	// 上传流量
	D                int64  `json:"d"`                  	// 下载流量
	TransferEnable   int64  `json:"transfer_enable"`    	// 套餐总流量
	Banned           bool   `json:"banned"`             	// 是否被封禁
	IsAdmin          bool   `json:"is_admin"`           	// 是否为管理员
	LastLoginAt      int    `json:"last_login_at"`      	// 最后登录时间
	IsStaff          bool   `json:"is_staff"`           	// 是否为员工
	LastLoginIP      string `json:"last_login_ip"`      	// 最后登录IP
	UUID             string `json:"uuid"`               	// 用户 UUID
	GroupID          int    `json:"group_id"`           	// 权限组 ID
	PlanID           int    `json:"plan_id"`            	// 订阅套餐
	SpeedLimit       int    `json:"speed_limit"`        	// 速度限制
	RemindExpire     bool   `json:"remind_expire"`      	// 是否提醒到期
	RemindTraffic    bool   `json:"remind_traffic"`     	// 是否提醒流量
	Token            string `json:"token"`              	// API 令牌
	ExpiredAt        int64  `json:"expired_at"`         	// 到期时间
	DeviceLimit      int    `json:"device_limit"`       	// 设备限制
	OnlineCount      int    `json:"online_count"`       	// 在线设备数
	LastOnlineAt     time.Time `json:"last_online_at"`      // 最后在线时间
	Remarks          string `json:"remarks"`            	// 备注
	CreatedAt        int64  `json:"created_at"`         	// 创建时间
	UpdatedAt        int64  `json:"updated_at"`         	// 更新时间
}