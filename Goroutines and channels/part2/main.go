package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go randomNumber(ch)
	go randomNumber(ch)
	fmt.Println("Рандомное число:", <-ch)
	fmt.Println("Рандомное число:", <-ch)
	time.Sleep(time.Millisecond)
}

func randomNumber(ch chan int) {
	n := rand.Intn(100)
	ch <- n
	fmt.Printf("Число %d было передано по каналу\n", n)
}
