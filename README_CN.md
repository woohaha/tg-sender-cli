# tg-sender-cli

[English](README.md)

一个简单的命令行工具，用于发送文件到 Telegram。

## 功能

- 智能识别文件类型（图片/视频/文档）
- 支持自定义配置文件
- 可选消息说明

## 安装

### 从源码编译

```bash
git clone git@github.com:woohaha/tg-sender-cli.git
cd tg-sender-cli
go build -o tg-sender .
cp tg-sender ~/bin/
```

### 从 Release 下载

前往 [Releases](https://github.com/woohaha/tg-sender-cli/releases) 下载对应平台的二进制文件。

### 前置条件

- Go 1.21+（仅编译需要）
- Telegram Bot Token（从 [@BotFather](https://t.me/BotFather) 获取）
- Chat ID（群组或频道的 ID）

## 配置

在 `~/.config/tg-sender/config.toml` 创建配置文件：

```bash
mkdir -p ~/.config/tg-sender
cat > ~/.config/tg-sender/config.toml << 'EOF'
bot_token = "YOUR_BOT_TOKEN"
chat_id = YOUR_CHAT_ID
EOF
```

## 使用方法

```bash
# 发送文件
tg-sender -f /path/to/file.jpg

# 带说明发送
tg-sender -f /path/to/file.jpg -m "你好世界"

# 使用自定义配置
tg-sender -f /path/to/file.jpg -c /path/to/config.toml
```

### 参数说明

| 参数 | 说明 |
|------|------|
| `-f` | 要上传的文件路径（必填） |
| `-m` | 消息说明（可选） |
| `-c` | 自定义配置文件路径（可选） |
| `-h` | 显示帮助信息 |

### 文件类型处理

| 类型 | 扩展名 | Telegram API |
|------|--------|--------------|
| 图片 | jpg, jpeg, png, gif, webp | sendPhoto |
| 视频 | mp4, mov, avi, mkv | sendVideo |
| 文档 | 其他所有文件 | sendDocument |

## 开发

### 项目结构

```
tg-sender-cli/
├── main.go           # 入口，命令行解析
├── config/
│   ├── config.go     # 配置加载
│   └── config_test.go
├── sender/
│   ├── sender.go     # Telegram 发送
│   └── sender_test.go
├── go.mod
└── go.sum
```

### 编译

```bash
go build -o tg-sender .
```

### 测试

```bash
go test ./...
```

### 依赖

- [BurntSushi/toml](https://github.com/BurntSushi/toml) - TOML 解析
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram Bot API

## 许可证

Apache 2.0
