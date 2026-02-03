# tg-sender-cli

[中文文档](README_CN.md)

A simple CLI tool to send files to Telegram.

## Features

- Smart file type detection (photo/video/document)
- Custom config file support
- Optional message caption

## Installation

### From source

```bash
git clone git@github.com:woohaha/tg-sender-cli.git
cd tg-sender-cli
go build -o tg-sender .
cp tg-sender ~/bin/
```

### Prerequisites

- Go 1.21+
- A Telegram Bot Token (get from [@BotFather](https://t.me/BotFather))
- A Chat ID (can be a group or channel)

## Configuration

Create config file at `~/.config/tg-sender/config.toml`:

```bash
mkdir -p ~/.config/tg-sender
cat > ~/.config/tg-sender/config.toml << 'EOF'
bot_token = "YOUR_BOT_TOKEN"
chat_id = YOUR_CHAT_ID
EOF
```

## Usage

```bash
# Send a file
tg-sender -f /path/to/file.jpg

# Send with caption
tg-sender -f /path/to/file.jpg -m "Hello world"

# Use custom config
tg-sender -f /path/to/file.jpg -c /path/to/config.toml
```

### Flags

| Flag | Description |
|------|-------------|
| `-f` | File path to upload (required) |
| `-m` | Message caption (optional) |
| `-c` | Custom config file path (optional) |

### File Type Handling

| Type | Extensions | Telegram API |
|------|------------|--------------|
| Photo | jpg, jpeg, png, gif, webp | sendPhoto |
| Video | mp4, mov, avi, mkv | sendVideo |
| Document | all others | sendDocument |

## Development

### Project Structure

```
tg-sender-cli/
├── main.go           # Entry point, CLI parsing
├── config/
│   ├── config.go     # Config loading
│   └── config_test.go
├── sender/
│   ├── sender.go     # Telegram sender
│   └── sender_test.go
├── go.mod
└── go.sum
```

### Build

```bash
go build -o tg-sender .
```

### Test

```bash
go test ./...
```

### Dependencies

- [BurntSushi/toml](https://github.com/BurntSushi/toml) - TOML parser
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram Bot API

## License

Apache 2.0
