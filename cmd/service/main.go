package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof" // Важливо для Lab #3
	"os"
	"time"

	"lab4-team-dev/internal/processor"

	"github.com/rs/zerolog"
)

func main() {
	// 1. Налаштовуємо прапорці
	workers := flag.Int("workers", 5, "number of worker goroutines")
	reportEvery := flag.Duration("report-every", 2*time.Second, "interval for logging processed stats")
	flag.Parse()

	// 2. Ініціалізуємо структурований логер zerolog
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "image-metadata-processor").Logger()

	// 3. Запускаємо pprof на окремому порті
	go func() {
		logger.Info().Msg("Pprof server started on :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			logger.Error().Err(err).Msg("Pprof server error")
		}
	}()

	// 4. Логуємо старт сервісу з параметрами
	logger.Info().Int("workers", *workers).Dur("report_every", *reportEvery).Msg("Image Metadata Processor started...")

	// 5. Запускаємо воркер-пул в окремій горутині,
	// щоб основна програма могла йти далі і виводити логи
	go processor.RunWorkerPool(*workers)

	// 6. Запускаємо таймер для періодичного звіту
	ticker := time.NewTicker(*reportEvery)
	defer ticker.Stop()

	// Цей цикл буде працювати вічно і виводити лог кожні 2 секунди
	for range ticker.C {
		logger.Info().
			Int("total_processed", 0). // 0, оскільки немає доступу до реальної статистики пулу
			Msg("processing stats")
	}
}
