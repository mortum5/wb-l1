package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * Реализовать все возможные способы остановки выполнения горутины.
 */

// Для контекста есть несколько вариантов WithCancel, WithTimeout, WithDeadline -- но по сути механизм один и тотже.
// Также здесь можно использовать select, либо for.
func cancelWithContext(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("cancel with context after 5 second")
}

// Здесь в зависимости от ситуации можно делать явную проверку на закрытие канала, если канал использовался для передачи значений
// или он буфферизированный.
// А можно просто читать значение из небуфферизированного.
func cancelWithChannel(done chan struct{}) {
	<-done
	fmt.Println("cancel after channel close")
}

func cancelWithVariable(isStop *bool) {
	var count int = 0
	for {
		if *isStop {
			break
		}

		time.Sleep(time.Millisecond * 200)
		count++
	}
	fmt.Println("cancel after shared variable has changed")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ch := make(chan struct{})
	var isStop bool = false

	go cancelWithContext(ctx)
	go cancelWithChannel(ch)
	go cancelWithVariable(&isStop)

	time.Sleep(time.Second * 5)
	isStop = true
	close(ch)

	time.Sleep(time.Second)

}
