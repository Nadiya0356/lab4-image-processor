package processor

import (
	"testing"
)

func TestNewProcessor(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if p == nil {
		t.Fatal("expected processor to be non-nil")
	}
}

func TestProcessImage_ValidFile(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("failed to create processor: %v", err)
	}

	err = p.ProcessImage("photo.jpg")
	if err != nil {
		t.Errorf("expected no error for valid file, got: %v", err)
	}

	// Перевіряємо що лічильник збільшився
	if p.GetProcessedCount() != 1 {
		t.Errorf("expected 1 processed image, got %d", p.GetProcessedCount())
	}
}

func TestProcessImage_InvalidFormat(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("failed to create processor: %v", err)
	}

	err = p.ProcessImage("document.pdf")
	if err == nil {
		t.Error("expected error for unsupported format, got nil")
	}

	// Лічильник не повинен збільшитись
	if p.GetProcessedCount() != 0 {
		t.Errorf("expected 0 processed images, got %d", p.GetProcessedCount())
	}
}

func TestProcessImage_Counter(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("failed to create processor: %v", err)
	}

	files := []string{"img1.jpg", "img2.png", "img3.webp"}
	for _, f := range files {
		p.ProcessImage(f)
	}

	if p.GetProcessedCount() != 3 {
		t.Errorf("expected 3 processed images, got %d", p.GetProcessedCount())
	}
}
