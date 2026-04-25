package stats

import (
	"sync"
	"testing"
)

// TestRaceConditionDemo навмисно запускає гонитву даних (Data Race).
func TestRaceConditionDemo(t *testing.T) {
	var wg sync.WaitGroup
	numGoroutines := 200

	// Запускаємо 200 горутин, які одночасно викликають функцію запису
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			IncrementProcessed("jpeg")
		}()
	}

	wg.Wait() // Чекаємо завершення всіх горутин
}

// TestIncrementProcessedConcurrent для перевірки після виправлення
func TestIncrementProcessedConcurrent(t *testing.T) {
	// Очищаємо мапу перед тестом (щоб тести не впливали один на одного)
	GlobalStats = make(map[string]int)

	var wg sync.WaitGroup
	numGoroutines := 200

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			IncrementProcessed("png")
		}()
	}

	wg.Wait()

	// Перевіряємо, чи правильний результат після безпечного конкурентного запису
	if GlobalStats["png"] != numGoroutines {
		t.Errorf("Очікувалося %d, отримано %d", numGoroutines, GlobalStats["png"])
	}
}
