package main

import (
	"fmt"
)

/**
 * Реализовать бинарный поиск встроенными методами языка.
 */

func search(a []int, x int) (int, bool) {
	n := len(a)

	left, right := 0, n

	for left < right {
		mid := left + (right-left)/2

		if a[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left, right < n && a[left] == x

}

func main() {
	arr := []int{1, 5, 10, 12, 33, 39, 44, 45, 46, 46, 47, 47, 47, 80, 81}
	show(search(arr, -1))
	show(search(arr, 1))
	show(search(arr, 5))
	show(search(arr, 43))
	show(search(arr, 44))
	show(search(arr, 46))
	show(search(arr, 47))
	show(search(arr, 81))
	show(search(arr, 90))

}

func show(x int, b bool) {
	if b {
		fmt.Println("Found at ", x)
	} else {
		fmt.Println("Not found")
	}
}
