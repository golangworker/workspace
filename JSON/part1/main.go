package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Age       int     `json:"age"`
	IsMarried bool    `json:"ismarried"`
	Height    float64 `json:"height"`
}

func (u User) Println() {
	fmt.Printf("Name: %v\n", u.Name)
	fmt.Printf("Address: %v\n", u.Address)
	fmt.Printf("Age: %v y.o.\n", u.Age)
	fmt.Printf("Is married: %v\n", u.IsMarried)
	fmt.Printf("Height: %.1f sm\n", u.Height)
}

var user = User{"Vladimir", "Moscow", 18, false, 182.0}

func userJSON(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("err:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Println()
	fmt.Fprintln(w, "success! server read your data")
}

func serverJSON(w http.ResponseWriter, r *http.Request) {
	if b, err := json.Marshal(user); err != nil {
		log.Println("err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(b)
		log.Println("success! user get server json")
	}
}

func main() {
	http.HandleFunc("/user", userJSON)
	http.HandleFunc("/server", serverJSON)

	log.Println("server has running in http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Println("err:", err.Error())
	}
}
