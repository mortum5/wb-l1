package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	m := make(map[rune]struct{})
	lowerStr := strings.ToLower(s)
	for _, c := range lowerStr {
		_, ok := m[c]
		if !ok {
			m[c] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "exuzjU"
	fmt.Print(isUnique(s))
}
