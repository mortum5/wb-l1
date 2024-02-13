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
	wg     sync.WaitGroup
	result chan int = make(chan int, N)
)

func incrementator() {
	var counter int = 0
	for i := 0; i < M; i++ {
		counter++
	}
	result <- counter
	wg.Done()
}

func main() {

	for i := 0; i < N; i++ {
		wg.Add(1)
		go incrementator()
	}

	wg.Wait()
	close(result)
	sum := 0
	for v := range result {
		sum += v
	}

	fmt.Println(sum)
}
