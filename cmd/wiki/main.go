package main

import (
	"gowiki/internal/logger"
	"gowiki/internal/wiki"
	"log"
	"net/http"
)

func main() {
	logger.Init()

	logger.Info.Println("Сервер запускается...")

	wiki.RegisterHandlers()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error.Printf("Ошибка запуска сервера: %v", err)
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
