package processor

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var LeakCache = make(map[string][]byte)

// ❗ компілюємо regex один раз (оптимізація CPU)
var imageRegex = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(20 * time.Millisecond)
			}
		}(i)
	}

	select {}
}

func processImage(workerID int) {
	var builder strings.Builder
	builder.WriteString("image_data_")
	builder.WriteString(fmt.Sprint(workerID))
	builder.WriteString("_timestamp_")
	builder.WriteString(fmt.Sprint(time.Now().UnixNano()))

	data := builder.String()

	if imageRegex.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		// ❗ FIX memory leak (обмеження кешу)
		if len(LeakCache) > 1000 {
			for k := range LeakCache {
				delete(LeakCache, k)
				break
			}
		}

		LeakCache[key] = make([]byte, 1024*10)
	}
}
