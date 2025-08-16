package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type candidates struct {
	Ivan atomic.Int64
	Dmitry atomic.Int64
	Lena atomic.Int64
}

func main() {
	Candidates := &candidates{}
	wg := &sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go voteForIvan(wg, Candidates)
	}
	wg.Wait()
	fmt.Println("Количество голосов за Ивана:", Candidates.Ivan.Load())

}

func voteForIvan(wg *sync.WaitGroup, c *candidates) {
	defer wg.Done()
	c.Ivan.Add(1)
}
