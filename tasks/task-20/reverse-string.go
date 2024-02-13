package main

import (
	"fmt"
	"strings"
)

/**
 * Разработать программу, которая переворачивает слова в строке.
 * Пример: «snow dog sun — sun dog snow».
 */

func reverse(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < (len(words)-1)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(reverse(s))
}
