package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("ошибка, запись в nil-мапу")
			fmt.Println(p)
		}
	}()
	
	var m map[int]any // nil-мапа
	m[1] = "попытка записи в nil-мапу"
	fmt.Println(m)
}
