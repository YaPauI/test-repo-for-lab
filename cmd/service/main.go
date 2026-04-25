package main

import (
	"lab3/internal/processor"
	"log"
	"net/http"
	_ "net/http/pprof" // Важливо для Lab #3
)

func main() {
	// Запускаємо pprof на окремому порті
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060",
			nil))
	}()
	log.Println("Image Metadata Processor started...")
	// Запускаємо "проблемний" обробник
	processor.RunWorkerPool(5)
}
