package main

import (
	"fmt"
)

/**
 * Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
 * во второй — результат операции x*2, после чего данные из второго канала
 * должны выводиться в stdout.
 */

const (
	N int = 5
)

func main() {
	arr := []int{7, 5, 4, 3, 2, 1, 8, 9, 78, -5}

	fCh := make(chan int, N)
	sCh := make(chan int, N)

	go func(in chan int) {
		for _, v := range arr {
			in <- v
		}
		close(in)
	}(fCh)

	go func(in, out chan int) {
		for v := range in {
			out <- v * v
		}
		close(sCh)
	}(fCh, sCh)

	for v := range sCh {
		fmt.Print(v, " ")
	}

}
