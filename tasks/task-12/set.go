package main

import "fmt"

/**
 * Имеется последовательность строк - (cat, cat, dog, cat, tree)
 * создать для нее собственное множество.
 */

type Set map[string]struct{}

func NewSet(arr []string) (s Set) {
	s = make(Set)
	for _, v := range arr {
		s[v] = struct{}{}
	}
	return
}

func main() {
	s := []string{"cat", "dog", "cat", "cat", "tree"}

	set := NewSet(s)
	fmt.Print(set)
}
