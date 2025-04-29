package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	logger "gowiki/internal/logger"
	"os"
)

var DB *sql.DB

func Init() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "go_wiki"),
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		logger.Error.Fatalf("Ошибка при открытии соединения с БД: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		logger.Error.Fatalf("БД не отвечает: %v", err)
	}

	logger.Info.Println("Успешное подключение к БД")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
