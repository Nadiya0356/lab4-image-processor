package stats

import "sync"

var (
	GlobalStats = make(map[string]int)
	mu          sync.RWMutex
)

func IncrementProcessed(imageType string) {
	mu.Lock()
	defer mu.Unlock()
	GlobalStats[imageType]++
}

func GetStats() map[string]int {
	mu.RLock()
	defer mu.RUnlock()

	copy := make(map[string]int)
	for k, v := range GlobalStats {
		copy[k] = v
	}
	return copy
}
