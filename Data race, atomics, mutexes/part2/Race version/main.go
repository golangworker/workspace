package main

import (
	"fmt"
	"sync"
)

var postman []string

func main() {
	wg := &sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go messageFan(wg)
	}
	wg.Wait()
	fmt.Println("Количество новых писем:", len(postman))
}

func messageFan(wg *sync.WaitGroup) {
	defer wg.Done()
	postman = append(postman, "Привет, я твой подписчек!")
}

