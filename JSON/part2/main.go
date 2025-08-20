package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

type Message struct {
	Title    string `json:"title"`
	Postcode int    `json:"postcode"`
	Text     string `json:"text"`
	IsUrgent bool   `json:"isurgent"`
}

type ID struct {
	V int `json:"id"`
}

var (
	allMessages map[int]Message = make(map[int]Message)
	mtx         sync.RWMutex
)

// хендлер для добавления нового сообщения пользователя и
// вывода всех сообщений, вместе с новым
func addMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("пользователь попытался добавить новое сообщение с неверным методом", r.Method)
		return
	}
	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("ошибка в чтении тела запроса:", err)
		return
	}
	log.Printf("пользователь %v ввел новое сообщение: %v", m.Title, m.Text)

	n := rand.Intn(1000) + 1
	mtx.Lock()
	allMessages[n] = m
	if err := json.NewEncoder(w).Encode(allMessages); err != nil {
		log.Println("ошибка в чтении json:", err)
		return
	}
	mtx.Unlock()
}

// хендлер для удаления сообщения по его ID, и вывод
// всех сообщений, без старого
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("пользователь попытался удалить сообщение с неверным методом", r.Method)
		return
	}
	var id ID
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		log.Println("ошибка в чтении json:", err)
		return
	}
	if message, ok := allMessages[id.V]; !ok {
		log.Println("ошибка, пользователь ввел несуществующий ID")
		return
	} else {
		log.Printf("пользователь ввел ID: %d для удаления сообщения: %s\n", id, message.Text)
	}
	mtx.Lock()
	delete(allMessages, id.V)
	for _, v := range allMessages {
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Println("ошибка в чтении json:", err)
			return
		}
	}
	mtx.Unlock()
}

// хендлер для получения всего списка сообщений
func showMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("пользователь попытался посмотреть список сообщений с неверным методом", r.Method)
		return
	}
	mtx.RLock()
	if err := json.NewEncoder(w).Encode(allMessages); err != nil {
		log.Println("ошибка в чтении json:", err)
		return
	}
	mtx.RUnlock()
}

// хендлер для обработки ID и вывода в теле ответа сообщение,
// принадлежащее ему, иначе статус nod found
func showMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("пользователь попытался посмотреть сообщение с неверным методом", r.Method)
		return
	}
	var id ID
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		log.Println("ошибка в чтении json:", err)
		return
	}
	if m, ok := allMessages[id.V]; !ok {
		log.Println("ошибка, пользователь ввел несуществующий ID")
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		log.Printf("пользователь ввел ID: %d для просмотра сообщения: %s\n", id.V, m.Text)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			log.Println("ошибка в чтении json:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/add", addMessage)
	http.HandleFunc("/delete", deleteMessage)
	http.HandleFunc("/list", showMessages)
	http.HandleFunc("/message", showMessage)

	log.Println("сервер запущен на http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера", err)
		return
	}
}
