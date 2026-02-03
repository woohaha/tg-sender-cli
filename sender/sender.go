package sender

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FileType int

const (
	FileTypeDocument FileType = iota
	FileTypePhoto
	FileTypeVideo
)

var photoExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
}

var videoExts = map[string]bool{
	".mp4": true, ".mov": true, ".avi": true, ".mkv": true,
}

func DetectFileType(filename string) FileType {
	ext := strings.ToLower(filepath.Ext(filename))
	if photoExts[ext] {
		return FileTypePhoto
	}
	if videoExts[ext] {
		return FileTypeVideo
	}
	return FileTypeDocument
}

type Sender struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func New(botToken string, chatID int64) (*Sender, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}
	return &Sender{bot: bot, chatID: chatID}, nil
}

func (s *Sender) SendFile(filePath string, caption string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filePath)
	}

	fileType := DetectFileType(filePath)

	var msg tgbotapi.Chattable
	switch fileType {
	case FileTypePhoto:
		photo := tgbotapi.NewPhoto(s.chatID, tgbotapi.FilePath(filePath))
		photo.Caption = caption
		msg = photo
	case FileTypeVideo:
		video := tgbotapi.NewVideo(s.chatID, tgbotapi.FilePath(filePath))
		video.Caption = caption
		msg = video
	default:
		doc := tgbotapi.NewDocument(s.chatID, tgbotapi.FilePath(filePath))
		doc.Caption = caption
		msg = doc
	}

	_, err := s.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send: %w", err)
	}
	return nil
}
