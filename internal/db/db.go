package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	connStr := "user=postgres password=postgres dbname=go_wiki sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения с БД: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("БД не отвечает: %v", err)
	}

	fmt.Println("Успешное подключение к БД")
}
