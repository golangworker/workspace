package main

import (
	"fmt"
	"time"
)

func main() {
	go iAmGoroutine(1)
	go iAmGoroutine(2)
	go iAmGoroutine(3)
	// ставим таймер для точного выполнения всех горутин
	time.Sleep(5 * time.Second)

}

func iAmGoroutine(n int) {
	for x := range 5 {
		fmt.Printf("Я горутина %d, делаю вывод на экран в %x раз\n", n, x)
		time.Sleep(time.Second)
	}
}
