package main

import (
	"gowiki/internal/wiki"
	"log"
	"net/http"
)

func main() {
	wiki.RegisterHandlers()
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
