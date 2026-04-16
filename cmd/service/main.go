package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	// Ініціалізація логера
	logger := log.With().Str("service", "image-processor").Logger()

	logger.Info().Msg("Service started")

	processedCount := 0

	// Імітація обробки зображень
	for i := 0; i < 3; i++ {
		logger.Info().Msg("Processing image")

		processedCount++

		logger.Info().
			Int("processed_images", processedCount).
			Msg("Images processed")
	}

	fmt.Println("Done")
}
