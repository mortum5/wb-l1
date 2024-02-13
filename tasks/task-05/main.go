package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
 * Разработать программу, которая будет последовательно отправлять значения в канал,
 * а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
 */

const (
	N           = 5
	BUFFER_SIZE = 1000
)

func producer(ctx context.Context, out chan<- int) {
	var counter int = 0
	timer := time.NewTicker(time.Millisecond * time.Duration(100))
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%d was sended\n", counter)
			fmt.Println("close producer")
			close(out)
			timer.Stop()
			return
		case <-timer.C:
			r := rand.Intn(9999) + 1
			counter++
			out <- r
		}
	}
}

func consumer(in <-chan int, exit chan<- struct{}) {
	var count int = 0
	for range in {
		count++
	}
	fmt.Println("cosumer received ", count, " value")
	close(exit)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*N)
	defer cancel()

	ch := make(chan int, BUFFER_SIZE)
	exit := make(chan struct{})

	go producer(ctx, ch)
	time.Sleep(1 * time.Second)
	go consumer(ch, exit)

	<-exit
}
