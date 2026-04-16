package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Str("service", "image-processor").Logger()

	logger.Info().Msg("Service started")

	processedCount := 0

	for i := 0; i < 3; i++ {
		logger.Info().Msg("Processing image")

		processedCount++

		logger.Info().
			Int("processed_images", processedCount).
			Msg("Images processed")
	}

	fmt.Println("Done")
}
