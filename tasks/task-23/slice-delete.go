package main

import "fmt"

var (
	N int = 15
)

func newSlice() []int {
	return make([]int, N)
}

func deleteIth(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func anotherDeleteIth(arr []int, i int) []int {
	newArr := make([]int, len(arr)-1)
	copy(newArr, append(arr[:i], arr[i+1:]...))
	return newArr
}

func main() {
	arr := newSlice()
	arr[5] = 1
	fmt.Println(len(arr), cap(arr), arr)
	arr = deleteIth(arr, 5)
	fmt.Println(len(arr), cap(arr), arr)

	arr = newSlice()
	arr[5] = 1
	fmt.Println(len(arr), cap(arr), arr)
	arr = anotherDeleteIth(arr, 5)
	fmt.Println(len(arr), cap(arr), arr)

}
