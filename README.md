# XBoard 机器人

## 项目介绍
XBoard 机器人是一个基于 Go 语言和 telebot.v4 库开发的 Telegram 机器人，旨在为用户提供与 XBoard 平台的便捷交互。通过简单的命令，用户可以绑定 XBoard 账户、查看账户信息、获取订阅链接、检查流量使用情况、执行每日签到等操作。项目采用模块化设计，易于维护和扩展，适合需要自动化管理 XBoard 账户的场景。

## 功能特性
- 将 Telegram 用户与 XBoard 账户绑定
- 查看用户信息
- 获取订阅链接
- 检查流量使用情况
- 执行每日签到
- 关于和帮助命令

## 安装设置
1. 克隆仓库：
   ```bash
   git clone https://github.com/rebe554owen/xboard-bot.git
   cd xboard-bot
   ```
2. 复制示例配置文件并编辑：
   ```bash
   cp config.example.yaml config.yaml
   ```
3. 打开 `config.yaml` 文件，填写 Telegram 机器人令牌和 XBoard API 相关信息：
   ```yaml
   telegram:
     token: "你的 Telegram 机器人令牌"
   xboard:
     address: "https://你的 XBoard API 地址"
     token: "你的 XBoard API 令牌"
   ```
4. 安装依赖：
   ```bash
   go mod tidy
   ```
5. 运行机器人：
   ```bash
   go run main.go
   ```

# 项目结构
   ```
   xboard-bot/
   ├── bot/                        # Telegram 机器人核心模块
   │   ├── bot.go                  # 机器人初始化和启动逻辑
   │   ├── handlers/               # 命令处理模块
   │   │   └── handlers.go         # 处理逻辑
   │   └── middleware/             # 中间件模块
   │       └── middleware.go       # 日志记录
   ├── config/                     # 配置加载模块
   │   └── config.go               # 使用 viper 解析 config.yaml 文件
       ├── config.example.yaml         # 示例配置文件，需复制为 config.yaml 使用
   ├── xboard/                     # XBoard API 客户端模块
   │   ├── client.go               # XBoard API 请求处理
   │   └── types.go                # XBoard 数据结构定义
   ├── main.go                     # 程序主入口，加载配置并启动机器人
   ├── go.mod                      # Go 模块定义，包含依赖信息
   └── README.md                   # 项目说明文档
   ```
## 许可证
本项目采用 [MIT 许可证](https://opensource.org/licenses/MIT)。
