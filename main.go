package main

import (
	"log"
	"net/http"
	"url-shortener/handler"
	"url-shortener/service"
	"url-shortener/storage"
)

func main() {
	storage := storage.NewMemoryStorage()
	service := service.NewURLService(storage)
	handler := handler.NewHandler(service)

	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/", handler.Redirect)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
