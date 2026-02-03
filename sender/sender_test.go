package sender

import "testing"

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
