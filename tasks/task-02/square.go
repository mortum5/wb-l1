package main

import (
	"fmt"
	"sync"
)

/**
 * Написать программу, которая конкурентно рассчитает значение квадратов чисел
 * взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
 */

func main() {
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Add(1)

		go func(i int) {
			fmt.Print(i*i, " ")
			wg.Done()
		}(v)
	}

	wg.Wait()

}
