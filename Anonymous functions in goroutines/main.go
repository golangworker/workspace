package main

import (
	"fmt"
	"time"
)

func main() {
	go func (n int) {
		for x:= range 3 {
			fmt.Printf("Привет №%d от %d горутины!\n", x, n)
		}
	}(1)

	go func () {
		for range 3 {
			fmt.Println("Как дела???")
		}
	}()

	go func () {
		for range 3 {
			fmt.Println("ха-ха-ха")
		}
	}()
	time.Sleep(time.Second)
}
