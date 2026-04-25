package main

import (
	"lab4-team-dev/internal/processor" // або "lab3/internal/processor", залежить як у тебе названо модуль
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Змінюємо цей лог, щоб створити конфлікт
	log.Println("Image Metadata Processor started with 10 WORKERS...")

	// Змінюємо кількість воркерів з 5 на 10
	processor.RunWorkerPool(10)
}
