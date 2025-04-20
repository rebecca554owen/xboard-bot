package xboard

// User 表示 XBoard 用户
type User struct {
	Email      string  `json:"email"`      // 邮箱
	Balance    float64 `json:"balance"`    // 余额
	Commission float64 `json:"commission"` // 佣金
	PlanID     int     `json:"plan_id"`    // 计划 ID
}

// Traffic 表示流量信息
type Traffic struct {
	Upload    int64 `json:"upload"`    // 上传流量
	Download  int64 `json:"download"`  // 下载流量
	Total     int64 `json:"total"`     // 总流量
	Remaining int64 `json:"remaining"` // 剩余流量
}