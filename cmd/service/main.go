package main

import (
	"fmt"
	"lab3-detector/internal/processor"
)

func main() {
	p, err := processor.NewProcessor()
	if err != nil {
		fmt.Printf("Failed to initialize processor: %v\n", err)
		return
	}
	defer p.Shutdown()

	// Тестова обробка зображень
	images := []string{
		"photo1.jpg",
		"photo2.png",
		"document.pdf", // невалідний — буде попередження
		"photo3.webp",
	}

	for _, img := range images {
		if err := p.processImage(img); err != nil {
			fmt.Printf("Error processing %s: %v\n", img, err)
		}
	}

	fmt.Printf("\nTotal processed: %d images\n", p.GetProcessedCount())
}
