package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("пользователь с неверным методом", r.Method)
		return
	}
	w.Write([]byte("привет мир!"))
	log.Println("HTTP хендлер отработал успешно!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("сервер запущен на http://localhost:8080/")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Println("ошибка ротера", err)
	}
}
