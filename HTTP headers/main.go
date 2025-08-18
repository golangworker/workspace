package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	headerMap := r.Header
	name := r.Header.Get("Name")
	if name != "" {
		log.Println("пользователь ввел свое имя:", name)
		fmt.Fprintf(w, "Hello %s, nice to meet you!!\n", name)
	}

	for k, v := range headerMap {
		fmt.Fprintf(w, "Key: %s, Value: %v\n", k, v)
	}
}

func main() {
	http.HandleFunc("/headers", handler)
	log.Println("сервер запущен на http://localhost:8080/headers")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка:", err)
	}
}
