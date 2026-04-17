package processor

import (
	"fmt"
	"regexp"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
)

// ImageMetadata зберігає метадані зображення
type ImageMetadata struct {
	Path   string
	Format string
	Size   int64
}

// Processor обробляє зображення та веде структуроване логування
type Processor struct {
	logger         *zap.Logger
	processedCount int64
	mu             sync.Mutex
	results        map[string]ImageMetadata
	formatRegex    *regexp.Regexp // винесено за межі циклу (оптимізація з Лаб №3)
}

// NewProcessor створює новий екземпляр Processor
func NewProcessor() (*Processor, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	regex, err := regexp.Compile(`\.(jpg|jpeg|png|gif|bmp|webp)$`)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %w", err)
	}

	logger.Info("Processor initialized")

	return &Processor{
		logger:      logger,
		results:     make(map[string]ImageMetadata),
		formatRegex: regex,
	}, nil
}

// processImage обробляє одне зображення та логує результат
func (p *Processor) processImage(path string) error {
	// Перевірка формату файлу
	if !p.formatRegex.MatchString(path) {
		p.logger.Warn("unsupported file format",
			zap.String("path", path),
		)
		return fmt.Errorf("unsupported format: %s", path)
	}

	// Симуляція обробки зображення
	metadata := ImageMetadata{
		Path:   path,
		Format: p.formatRegex.FindString(path),
		Size:   1024, // у реальному коді — реальний розмір файлу
	}

	// Захист від race condition (з Лаб №3)
	p.mu.Lock()
	p.results[path] = metadata
	p.mu.Unlock()

	// Атомарне збільшення лічильника
	count := atomic.AddInt64(&p.processedCount, 1)

	// Структуроване логування
	p.logger.Info("image processed",
		zap.String("path", path),
		zap.String("format", metadata.Format),
		zap.Int64("size_bytes", metadata.Size),
		zap.Int64("total_processed", count),
	)

	return nil
}

// GetProcessedCount повертає кількість оброблених зображень
func (p *Processor) GetProcessedCount() int64 {
	return atomic.LoadInt64(&p.processedCount)
}

// Shutdown коректно завершує роботу logger
func (p *Processor) Shutdown() {
	p.logger.Info("Processor shutting down",
		zap.Int64("total_images_processed", p.GetProcessedCount()),
	)
	p.logger.Sync()
}
