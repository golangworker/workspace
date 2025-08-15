package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func (ch chan int) {
		t := rand.Intn(10)*100 + 100 // от 100ms до 1000ms
		time.Sleep(time.Millisecond * time.Duration(t))
		fmt.Printf("Проспали: %dms\n", t)
		ch <- t

	}(ch1)

	go func (ch chan int) {
		t := rand.Intn(10)*100 + 100 // от 100ms до 1000ms
		time.Sleep(time.Millisecond * time.Duration(t))
		fmt.Printf("Проспали: %dms\n", t)
		ch <- t
	}(ch2)

	time.Sleep(time.Millisecond * 500)
	var randTime int
	select {
	case randTime = <- ch1:
		fmt.Println("Выполнилась горутина 1, со значением:", randTime)
	case randTime = <- ch2:
		fmt.Println("Выполнилась горутина 2, со значением:", randTime)
	default:
		fmt.Println("Никакая из горутин не выполнилась :(")
	}
} 
