package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	n := rand.Intn(10) + 1
	fmt.Printf("Начинаем работу с %d огородниками\n", n)
	for i := range n {
		wg.Add(1)
		go gardener(wg, i + 1)
	}
	wg.Wait()
	fmt.Println("Все огородники закончили работу! Программа завершена.")
}

func gardener(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	fmt.Printf("Огородник %d начал удобрять и поливать свою область\n", n)
	randomSleepTime()
	fmt.Printf("Огородник %d закончил удобрять и поливать свою область\n", n)
}

func randomSleepTime() {
	n := rand.Intn(501) + 500
	time.Sleep(time.Millisecond * time.Duration(n))
}
