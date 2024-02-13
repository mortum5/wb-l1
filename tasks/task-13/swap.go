package main

import "fmt"

/**
 * Поменять местами два числа без создания временной переменной.
 */

func main() {
	left, right := 5, 6
	fmt.Println("before swap", left, right)
	left, right = right, left
	fmt.Println("after swap", left, right)
}
