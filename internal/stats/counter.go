package stats

import "sync"

var (
	GlobalStats = make(map[string]int)
	mu          sync.Mutex
)

func IncrementProcessed(imageType string) {
	mu.Lock()
	defer mu.Unlock()

	GlobalStats[imageType]++
}
