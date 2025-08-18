package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	allMessages map[int]string = make(map[int]string)
	mtx         sync.Mutex
)

func addMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка в чтении тела запроса:", err)
	}
	if len(body) == 0 {
		log.Println("пользователь не ввел данные")
		return
	} else {
		stringMessage := string(body)
		log.Println("пользователь ввел новое сообщение:", stringMessage)
		fmt.Fprint(w, "All messages:")
		mtx.Lock()
		n := rand.Intn(1000)
		allMessages[n] = stringMessage
		for k, v := range allMessages {
			fmt.Fprintf(w, "\nID: %d, message --> %s", k, v)
		}
		mtx.Unlock()
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка в чтении тела запроса:", err)
	}
	if len(body) == 0 {
		log.Println("пользователь не ввел данные")
		return
	} else {
		id, err := strconv.Atoi(string(body))
		if err != nil {
			log.Println("ошибка, пользователь ввел не ID")
			return
		}
		if message, ok := allMessages[id]; !ok {
			log.Println("ошибка, пользователь ввел несуществующий ID")
			return
		} else {
			log.Printf("пользователь ввел ID: %d для удаления сообщения: %s\n", id, message)
		}
		fmt.Fprint(w, "All messages:")
		mtx.Lock()
		delete(allMessages, id)
		for k, v := range allMessages {
			fmt.Fprintf(w, "\nID:%d, message --> %s", k, v)
		}
		mtx.Unlock()
	}
}

func main() {
	http.HandleFunc("/add", addMessage)
	http.HandleFunc("/delete", deleteMessage)

	log.Println("сервер запущен на http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера", err)
		return
	}
}
