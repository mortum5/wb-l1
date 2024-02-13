package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/**
 * Реализовать постоянную запись данных в канал (главный поток). Реализовать
 * набор из N воркеров, которые читают произвольные данные из канала и
 * выводят в stdout. Необходима возможность выбора количества воркеров при
 * старте.
 *
 * Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
 * способ завершения работы всех воркеров.
 */

const (
	N = 5
)

var (
	wg sync.WaitGroup
)

/**
 * Создаётся канал для оповещения о сигналах системы. При получении значения из канала
 * происходит закрытие канала с данными. Так же для того, чтобы должать завершения каждой
 * из горутин заводится waitGroup.
 */
func worker(id int, in <-chan int) {
	for v := range in {
		fmt.Println("worker ", id, "value ", v)
	}
	fmt.Println("closing worker ", id)
	wg.Done()
}

func main() {
	exit := make(chan os.Signal, 1)
	in := make(chan int, N)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(i+1, in)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-exit:
			close(in)
			wg.Wait()
			return
		case <-ticker.C:
			randInt := rand.Intn(9999)
			in <- randInt
		}
	}

}
