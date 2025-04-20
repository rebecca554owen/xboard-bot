package xboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Client 与 XBoard API 交互
type Client struct {
	baseURL string        // API 基础地址
	token   string        // API 令牌
	bindMap sync.Map      // 存储用户 ID 到 XBoard 密钥的映射
}

// NewClient 创建新的 XBoard 客户端
func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

// BindUser 将 Telegram 用户绑定到 XBoard 密钥
func (c *Client) BindUser(userID int64, key string) error {
	// 实际实现中，应通过 API 验证密钥
	// 这里通过内存存储模拟绑定
	c.bindMap.Store(userID, key)
	return nil
}

// IsUserBound 检查用户是否已绑定
func (c *Client) IsUserBound(userID int64) bool {
	_, exists := c.bindMap.Load(userID)
	return exists
}

// GetUserInfo 获取用户信息
func (c *Client) GetUserInfo(userID int64) (*User, error) {
	_, ok := c.bindMap.Load(userID)
	if !ok {
		return nil, fmt.Errorf("用户未绑定")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/api/v1/user/info", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetSubscription 获取订阅链接
func (c *Client) GetSubscription(userID int64) (string, error) {
	_, ok := c.bindMap.Load(userID)
	if !ok {
		return "", fmt.Errorf("用户未绑定")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/api/v1/user/getSubscribe", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.URL, nil
}

// GetTraffic 获取流量信息
func (c *Client) GetTraffic(userID int64) (*Traffic, error) {
	_, ok := c.bindMap.Load(userID)
	if !ok {
		return nil, fmt.Errorf("用户未绑定")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/api/v1/user/traffic", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var traffic Traffic
	if err := json.NewDecoder(resp.Body).Decode(&traffic); err != nil {
		return nil, err
	}
	return &traffic, nil
}

// CheckIn 执行每日签到
func (c *Client) CheckIn(userID int64) (string, error) {
	_, ok := c.bindMap.Load(userID)
	if !ok {
		return "", fmt.Errorf("用户未绑定")
	}

	req, err := http.NewRequest("POST", c.baseURL+"/api/v1/user/checkin", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Message, nil
}