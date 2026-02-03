package sender

import (
	"testing"
)

func TestValidateSendParams(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		message  string
		wantErr  bool
	}{
		{"file only", "/path/to/file.jpg", "", false},
		{"file and message", "/path/to/file.jpg", "caption", false},
		{"message only", "", "hello", false},
		{"neither", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSendParams(tt.file, tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSendParams(%q, %q) error = %v, wantErr %v", tt.file, tt.message, err, tt.wantErr)
			}
		})
	}
}

func TestDetectFileType(t *testing.T) {
	tests := []struct {
		filename string
		want     FileType
	}{
		{"photo.jpg", FileTypePhoto},
		{"photo.jpeg", FileTypePhoto},
		{"photo.png", FileTypePhoto},
		{"photo.gif", FileTypePhoto},
		{"photo.webp", FileTypePhoto},
		{"photo.JPG", FileTypePhoto},
		{"video.mp4", FileTypeVideo},
		{"video.mov", FileTypeVideo},
		{"video.avi", FileTypeVideo},
		{"video.mkv", FileTypeVideo},
		{"video.MP4", FileTypeVideo},
		{"doc.pdf", FileTypeDocument},
		{"doc.txt", FileTypeDocument},
		{"archive.zip", FileTypeDocument},
		{"noextension", FileTypeDocument},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got := DetectFileType(tt.filename)
			if got != tt.want {
				t.Errorf("DetectFileType(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
