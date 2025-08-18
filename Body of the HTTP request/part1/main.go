package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	mtx           sync.Mutex
	sliceMessages []string
)

func saveAndPrintMessages(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка чтения данных тела пользователя:", err.Error())
		return
	}
	if len(b) == 0 {
		log.Println("пользователь ничего не ввел в тело запроса")
	} else {
		stringMessage := string(b)
		log.Println("пользователь ввел новое сообщение:", stringMessage)
		mtx.Lock()
		sliceMessages = append(sliceMessages, stringMessage)
		mtx.Unlock()
	}
	w.Write([]byte("Все сообщения:"))
	for _, v := range sliceMessages {
		w.Write([]byte("\n" + v))
	}
}

func main() {
	http.HandleFunc("/message", saveAndPrintMessages)
	log.Println("сервер запущен и доступен на паттерне http://localhost:8080/message")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера", err.Error())
		return
	}
}
