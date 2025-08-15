package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("ошибка, деление на ноль")
		}
	}()
	a, b := 5, 0
	fmt.Println(a / b)
}
