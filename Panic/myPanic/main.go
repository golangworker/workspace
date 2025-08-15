package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	for range 10 {
		defer func() {
			p := recover()
			if p != nil {
				fmt.Println("о нет, ошибка!")
				fmt.Println(p)
			}
		}()
		n = rand.Intn(5)
		if n == 0 {
			panic("число равно 0")
		}
		fmt.Printf("Операция прошла успешно, число равно %d\n", n)
	}
}
