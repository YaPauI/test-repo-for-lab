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
	workers := flag.Int("workers", 5, "number of worker goroutines")
	reportInterval := flag.Duration("report-every", 2*time.Second, "interval for logging processed stats")
	flag.Parse()

	// Розбиваємо ланцюжок ініціалізації для кращої читабельності
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("service", "image-metadata-processor").
		Logger()

	go func() {
		logger.Info().Msg("Pprof server started on :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			logger.Error().Err(err).Msg("Pprof server error")
		}
	}()

	logger.Info().
		Int("workers", *workers).
		Dur("report_interval", *reportInterval).
		Msg("Image Metadata Processor started...")

	go processor.RunWorkerPool(*workers)

	ticker := time.NewTicker(*reportInterval)
	defer ticker.Stop()

	for range ticker.C {
		logger.Info().
			Int("total_processed", 0).
			Msg("processing stats")
	}
}
