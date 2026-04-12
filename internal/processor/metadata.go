package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

// 🔹 глобальний кеш
var LeakCache = make(map[string][]byte)

// 🔹 mutex для map (FIX race condition)
var mu sync.Mutex

// 🔹 оптимізація regex (FIX CPU)
var re = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	select {} // блокування програми
}

func processImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d",
		workerID, time.Now().UnixNano())

	if re.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		mu.Lock()

		// ✅ ВИПРАВЛЕННЯ ВИТОКУ
		if len(LeakCache) > 1000 {
			LeakCache = make(map[string][]byte)
		}

		LeakCache[key] = make([]byte, 1024*200)

		mu.Unlock()
	}
}
