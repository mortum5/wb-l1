package main

import "fmt"

/**
 * Реализовать пересечение двух неупорядоченных множеств.
 */

type Set map[int]struct{}

func NewSet(arr []int) (s Set) {
	s = make(Set)
	for _, v := range arr {
		s[v] = struct{}{}
	}
	return
}

func (s Set) Intersect(s1 Set) (s2 Set) {
	s2 = make(Set)
	for k := range s {
		if _, ok := s1[k]; ok {
			s2[k] = struct{}{}
		}
	}
	return
}

func main() {
	arr1 := []int{1, 5, 3, 2, 6, 5, 3, 7, 8}
	arr2 := []int{9, 5, 4, 4, 4, 15, 6, 15, 7, 99}

	s1 := NewSet(arr1)
	s2 := NewSet(arr2)

	fmt.Print(s1.Intersect(s2))

}
