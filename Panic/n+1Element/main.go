package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("ошибка, выход за пределы слайса")
		}
	}()
	slice := []int{1, 2, 3}
	fmt.Println(slice[4])
}
