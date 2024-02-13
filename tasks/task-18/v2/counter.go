package main

import (
	"fmt"
	"sync"
)

const (
	N int = 100
	M int = 1000
)

var (
	counter int = 0
	lock    sync.Mutex
	wg      sync.WaitGroup
)

func incrementator() {
	for i := 0; i < M; i++ {
		lock.Lock()
		counter++
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	for i := 0; i < N; i++ {
		wg.Add(1)
		go incrementator()
	}

	wg.Wait()
	fmt.Println(counter)
}
