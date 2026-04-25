package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

// Додаємо ліміт для кешу
const maxCacheSize = 500

var (
	// Додаємо м'ютекс для захисту мапи
	leakCacheMu sync.Mutex
	LeakCache   = make(map[string][]byte, maxCacheSize)
)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	select {}
}

var imageRe = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

func processImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())
	// O(n) без накладних витрат на компіляцію
	if imageRe.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		leakCacheMu.Lock()
		// Очищення при переповненні
		if len(LeakCache) >= maxCacheSize {
			LeakCache = make(map[string][]byte, maxCacheSize)
		}
		LeakCache[key] = make([]byte, 1024*10)
		leakCacheMu.Unlock()
	}
}
