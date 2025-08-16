package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var mtx sync.RWMutex
var storage map[string]string = make(map[string]string)

func main() {
	t := time.Now()
	wg := &sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		putInStorage(wg)
	}
	wg.Wait()
	fmt.Println("Прошло времени:", time.Since(t))
}

func putInStorage(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.Lock()
	storage[setKey()] = setValue()
	mtx.Unlock()
}

func setValue() string {
	var wordSlice []byte
	for range 5 {
		n := rand.Intn(35) + 65
		wordSlice = append(wordSlice, byte(n))
	}
	return string(wordSlice)
}

func setKey() string {
	n := int64(rand.Intn(100))
	return strconv.FormatInt(n, 10)
}

