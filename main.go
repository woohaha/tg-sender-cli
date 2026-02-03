package main

import (
	"flag"
	"fmt"
	"os"

	"tg-sender/config"
	"tg-sender/sender"
)

var version = "dev"

func main() {
	var (
		filePath    string
		message     string
		configPath  string
		showHelp    bool
		showVersion bool
	)

	flag.StringVar(&filePath, "f", "", "File path to upload (required)")
	flag.StringVar(&message, "m", "", "Message caption (optional)")
	flag.StringVar(&configPath, "c", "", "Config file path (optional, default: ~/.config/tg-sender/config.toml)")
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.BoolVar(&showVersion, "v", false, "Show version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `tg-sender - Send files to Telegram

USAGE:
    tg-sender -f <file> [-m <message>] [-c <config>]

FLAGS:
    -f <file>     File path to upload (required)
                  Supports: images (jpg/png/gif/webp), videos (mp4/mov/avi/mkv), documents (any)
    -m <message>  Message caption (optional)
    -c <config>   Config file path (optional)
                  Default: ~/.config/tg-sender/config.toml
    -h            Show this help message
    -v            Show version

EXAMPLES:
    # Send a photo
    tg-sender -f screenshot.png

    # Send with caption
    tg-sender -f report.pdf -m "Monthly report"

    # Use custom config
    tg-sender -f video.mp4 -c /path/to/config.toml

CONFIG FILE FORMAT (TOML):
    bot_token = "YOUR_BOT_TOKEN"
    chat_id = YOUR_CHAT_ID

FILE TYPE DETECTION:
    Photo:    .jpg, .jpeg, .png, .gif, .webp -> sendPhoto API
    Video:    .mp4, .mov, .avi, .mkv         -> sendVideo API
    Document: all other extensions           -> sendDocument API

EXIT CODES:
    0  Success
    1  Error (missing args, config error, send failed)

For AI Agents: This tool sends files to a preconfigured Telegram chat.
Required: -f flag with valid file path. Config must exist at default path
or specified via -c flag. Returns "File sent successfully" on success.
`)
	}

	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if showVersion {
		fmt.Printf("tg-sender version %s\n", version)
		os.Exit(0)
	}

	if filePath == "" {
		fmt.Fprintln(os.Stderr, "Error: -f (file path) is required")
		flag.Usage()
		os.Exit(1)
	}

	if configPath == "" {
		configPath = config.DefaultPath()
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config from %s: %v\n", configPath, err)
		fmt.Fprintln(os.Stderr, "\nCreate config file with:")
		fmt.Fprintf(os.Stderr, "  mkdir -p %s\n", config.DefaultPath()[:len(config.DefaultPath())-11])
		fmt.Fprintln(os.Stderr, "  cat > "+config.DefaultPath()+" << 'EOF'")
		fmt.Fprintln(os.Stderr, "bot_token = \"YOUR_BOT_TOKEN\"")
		fmt.Fprintln(os.Stderr, "chat_id = YOUR_CHAT_ID")
		fmt.Fprintln(os.Stderr, "EOF")
		os.Exit(1)
	}

	s, err := sender.New(cfg.BotToken, cfg.ChatID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating sender: %v\n", err)
		os.Exit(1)
	}

	if err := s.SendFile(filePath, message); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File sent successfully")
}
