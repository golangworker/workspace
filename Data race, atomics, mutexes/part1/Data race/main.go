package main

import (
	"fmt"
	"sync"
)

type candidates struct {
	Ivan int
	Lena int
	Dmitry int
}

func main() {
	wg := &sync.WaitGroup{}
	Candidates := &candidates{}
	for range 1000 {
		wg.Add(1)
		go voteForIvan(wg, Candidates)
	}
	wg.Wait()
	fmt.Println("Голосов за Ивана", Candidates.Ivan)
}

func voteForIvan(wg *sync.WaitGroup, c *candidates) {
	defer wg.Done()
	c.Ivan ++
}
