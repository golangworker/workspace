package main

import (
	"fmt"
	"time"
)

func main() {
	chString := make(chan string)
	chFloat := make(chan float64)
	chInt := make(chan int)

	go func() {
		for {
			time.Sleep(time.Millisecond * 300)
			chInt <- 42
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			chFloat <- 3.14
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			chString <- "Hello!"
		}
	}()

	// Запускаем цикл для получения нескольких значений
	for range 100 {
		select {
		case n := <- chString:
			fmt.Printf("String: %s\n", n)
		case n := <- chFloat:
			fmt.Printf("Float: %f\n", n)
		case n := <- chInt:
			fmt.Printf("Int: %d\n", n)
		}
	}
}
