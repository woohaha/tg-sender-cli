package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create temp config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.toml")

	content := `bot_token = "123456:ABC-DEF"
chat_id = -1001234567890
`
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.BotToken != "123456:ABC-DEF" {
		t.Errorf("BotToken = %q, want %q", cfg.BotToken, "123456:ABC-DEF")
	}

	if cfg.ChatID != -1001234567890 {
		t.Errorf("ChatID = %d, want %d", cfg.ChatID, -1001234567890)
	}
}

func TestLoadConfigFileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/config.toml")
	if err == nil {
		t.Error("Load() should return error for nonexistent file")
	}
}

func TestDefaultConfigPath(t *testing.T) {
	path := DefaultPath()
	if path == "" {
		t.Error("DefaultPath() should not return empty string")
	}

	// Should contain .config/tg-sender
	if !filepath.IsAbs(path) {
		t.Error("DefaultPath() should return absolute path")
	}
}
