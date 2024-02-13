package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"slices"
)

/**
 * Реализовать быструю сортировку массива встроенными методами языка
 */

func partition(arr []int, leftIndex, rightIndex int) int {
	pivot := arr[leftIndex+(rightIndex-leftIndex)/2]
	for leftIndex <= rightIndex {
		for arr[leftIndex] < pivot {
			leftIndex++
		}
		for pivot < arr[rightIndex] {
			rightIndex--
		}
		if leftIndex <= rightIndex {
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
			leftIndex++
			rightIndex--
		}
	}
	return leftIndex
}

func QuickSort(arr []int, leftBound, rightBound int) {
	if leftBound < rightBound {
		pivot := partition(arr, leftBound, rightBound)
		QuickSort(arr, leftBound, pivot-1)
		QuickSort(arr, pivot, rightBound)
	}
}

func main() {
	randomTest := randArray(77)

	randomTestAnswer := make([]int, len(randomTest))
	copy(randomTestAnswer, randomTest)

	slices.Sort(randomTestAnswer)

	testcases := []struct {
		Name   string
		Args   []int
		Result []int
	}{
		{
			Name:   "Empty",
			Args:   []int{},
			Result: []int{},
		},
		{
			Name:   "One elements",
			Args:   []int{1},
			Result: []int{1},
		},
		{
			Name:   "Already sorted",
			Args:   []int{-11, -5, 1, 5, 7, 9, 12},
			Result: []int{-11, -5, 1, 5, 7, 9, 12},
		},
		{
			Name:   "With duplicate elems",
			Args:   []int{-11, -11, 5, 5, 1, -20, 7, 7, 7, 9, 99, 1, 1, 1},
			Result: []int{-20, -11, -11, 1, 1, 1, 1, 5, 5, 7, 7, 7, 9, 99},
		},
		{
			Name:   "With random elems",
			Args:   randomTest,
			Result: randomTestAnswer,
		},
	}

	for _, tc := range testcases {
		QuickSort(tc.Args, 0, len(tc.Args)-1)
		if !reflect.DeepEqual(tc.Args, tc.Result) {
			fmt.Println("Error in ", tc.Name, " with result ", tc.Args)
		}
	}
}

func randArray(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, rand.Int())
	}
	return
}
