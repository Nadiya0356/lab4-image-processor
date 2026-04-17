package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"lab3-detector/internal/processor"
)

func main() {
	// pprof сервер
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Image Metadata Processor started...")

	// запуск воркерів
	processor.RunWorkerPool(5)
}
