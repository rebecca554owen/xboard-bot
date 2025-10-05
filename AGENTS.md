# Repository Guidelines

## 项目结构与模块组织
`main.go` 启动并加载配置；命令放在 `bot/handler`，中间件在 `bot/middleware`，共享适配器位于 `utils/`，配置解析由 `config/config.go` 读取根目录 `config.yaml`。扩展命令时在 `handler` 新建文件并统一注册。

## 构建、测试与开发命令
- `cp config.example.yaml config.yaml`：复制模板后补充密钥。
- `go run main.go`：读取本地配置启动。
- `go build ./...`：生成部署二进制。
- `go test ./...`：执行全部测试，可配 `-run` 聚焦。
- `go fmt ./...` 与 `golangci-lint run`：校验格式与静态规则。

## 代码风格与命名约定
遵循 `go fmt` 输出，tab 缩进。导出符号使用 PascalCase，私有符号用 camelCase；处理器以 `handleXxx` 命名，中间件以 `WithXxx` 命名；配置键保持小写短横线。

## 测试指南
新增逻辑需配套表驱动测试，命名 `Test<模块>_<场景>`。若依赖外部服务，通过接口注入 stub 或 mock；提交前执行 `go test -cover ./...` 并关注核心覆盖率。

## 提交与合并请求规范
提交信息保持简洁祈使句（如 `add bind handler`、`fix redis cache`），聚焦单一改动。PR 描述需说明背景、实现与验证，关联 issue，界面或行为变动附日志或截图。提交作者保持 `rebecaa554owen`，大改动按模块拆分。

## 配置与安全提示
敏感 token 只保存在 `config.yaml` 或环境变量，禁止提交仓库。部署前确认 Telegram 与 XBoard 接口使用 HTTPS 并可访问；新增外部依赖时更新文档并说明安全影响。
