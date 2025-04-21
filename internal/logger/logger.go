package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
)

func Init() {
	logFile, err := os.OpenFile("wiki.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Ошибка при создании файла логов: %v", err)
	}
	Info = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
}
