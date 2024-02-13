package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	N int = 100
	M int = 1000
)

var (
	counter atomic.Uint64
	wg      sync.WaitGroup
)

func incrementator() {
	for i := 0; i < 1000; i++ {
		counter.Add(1)
	}
	wg.Done()
}

func main() {
	for i := 0; i < N; i++ {
		wg.Add(1)
		go incrementator()
	}

	wg.Wait()
	fmt.Println(counter.Load())
}
