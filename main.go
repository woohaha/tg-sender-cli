package main

import (
	"flag"
	"fmt"
	"os"

	"tg-sender/config"
	"tg-sender/sender"
)

func main() {
	var (
		filePath   string
		message    string
		configPath string
	)

	flag.StringVar(&filePath, "f", "", "File path to upload (required)")
	flag.StringVar(&message, "m", "", "Message caption (optional)")
	flag.StringVar(&configPath, "c", "", "Config file path (optional)")
	flag.Parse()

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
