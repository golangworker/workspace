package main

import (
	"fmt"
	"net/http"
)

type Query struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	for k, v := range r.URL.Query() {
		fmt.Printf("key: %s value: %v", k, v)
	}
}

func main() {
	http.HandleFunc("/", queryHandler)
	fmt.Println("Сервер запущен на http://localhost:8080/")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("err:", err)
	}
}
