package main

import (
	"fmt"
	"sync"
)

var mtx sync.Mutex
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
	mtx.Lock()
	postman = append(postman, "Привет, я твой подписчек!")
	mtx.Unlock()
}
