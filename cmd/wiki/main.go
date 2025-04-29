package main

import (
	"gowiki/internal/db"
	handler "gowiki/internal/handler"
	"gowiki/internal/logger"
	"log"
	"net/http"
)

func main() {
	logger.Init()
	logger.Info.Println("Сервер запускается...")

	db.Init()

	handler.RegisterHandlers()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error.Printf("Ошибка запуска сервера: %v", err)
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
