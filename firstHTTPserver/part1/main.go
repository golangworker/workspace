package main

import (
	"log"
	"math/rand"
	"net/http"
)

var randomPhrases = []string{"hii!", "how are you?", "what is the weather like?", "How is your session?", "bye"}

func handler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(len(randomPhrases))
	randomPhrase := randomPhrases[n]
	w.Write([]byte(randomPhrase))
	log.Println("ответ пользователю отправлен")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("сервер запущен и доступен на http://localhost:8080/")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера", err.Error())
	}
}
