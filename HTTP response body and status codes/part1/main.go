package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
)

var badNums []int = []int{0, 6, 7, 8, 9}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка записи тела запроса: ", err)
		return
	}
	num, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println("введённые данные не являются числом")
		return
	}
	log.Println("пользователь ввел", num)
	lastDigit := num % 10
	if lastDigit < 0 {
		lastDigit = -lastDigit // Для отрицательных чисел
	}
	if slices.Contains(badNums, lastDigit) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "server not found 404")
		return
	}
	n := rand.Intn(100)
	statusCode := lastDigit*100 + n
	log.Println("новый статус код", statusCode)
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Nice? idk what you see...\n%d", statusCode)
}

func main() {
	http.HandleFunc("/statuscode", handler)
	log.Println("сервер запущен на http://localhost:8080/statuscode")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("ошибка ротера:", err)
	}
}
