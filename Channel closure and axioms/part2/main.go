package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	go randomSurvey(ch)
	for v := range ch {
		fmt.Println("Житель говорит:", v)
	}
	fmt.Println("Опрос завершён успешно!")
}

func randomSurvey(ch chan string) {
	n := rand.Intn(10) + 1
	for range n {
		time.Sleep(randomTime())
		ch <- randomString()
	}
	close(ch)

}

func randomString() string {
	var r byte
	rSlice := make([]byte, 0, 5) 
	for range 5 {
		alp := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"
		r = alp[rand.Intn(len(alp))]
		rSlice = append(rSlice, r)
	}
	return string(rSlice)
}

func randomTime() time.Duration {
	n := time.Duration(rand.Intn(401) + 300) * time.Millisecond
	return n
}
