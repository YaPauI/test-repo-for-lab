package stats

import "sync"

var (
	mu          sync.RWMutex
	GlobalStats = make(map[string]int)
)

func IncrementProcessed(imageType string) {
	mu.Lock()
	defer mu.Unlock()
	GlobalStats[imageType]++
}
