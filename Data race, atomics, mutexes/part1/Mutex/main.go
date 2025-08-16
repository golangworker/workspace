package main

import (
	"fmt"
	"sync"
)

var mtx sync.Mutex

type candidates struct {
	Ivan int
	Dmitry int
	Lena int
}

func main() {
	Candidates := &candidates{}
	wg := &sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go  voteForIvan(wg, Candidates)
	}
	wg.Wait()
	fmt.Println("Количество голосов за Ivan:", Candidates.Ivan)
}

func voteForIvan(wg *sync.WaitGroup, c *candidates) {
	defer wg.Done()
	mtx.Lock()
	c.Ivan ++
	mtx.Unlock()
}
