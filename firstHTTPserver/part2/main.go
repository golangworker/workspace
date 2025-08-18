package main

import (
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Я собака и я говорю Гав"))
	log.Println("Хендлер dog отправил ответ")
}

func cow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Я корова и я говорю Мууу"))
	log.Println("Хендлер cow отправил ответ")
}

func cat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Я кошка и я говорю Мяу"))
	log.Println("Хендлер cat отправил ответ")
}

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cow", cow)
	http.HandleFunc("/cat", cat)

	log.Println("сервер запущен на http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера", err.Error())
	}
}
