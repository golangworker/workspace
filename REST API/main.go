package main

import (
	"app/books"
	"app/http"
	"log"
)

func main() {
	library := books.NewLibrary()
	handlers := http.NewHTTPHandlers(library)
	httpServer := http.NewHTTPServer(handlers)

	if err := httpServer.Start(); err != nil {
		log.Println("Error starting HTTP server:", err)
	}
}
