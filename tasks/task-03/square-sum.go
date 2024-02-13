package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 * Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
 * с использованием конкурентных вычислений.
 */

func main() {
	var wg sync.WaitGroup
	var sum atomic.Int64
	arr := []int64{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Add(1)
		go func(i int64) {
			sum.Add(i * i)
			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Print(sum.Load())
}
