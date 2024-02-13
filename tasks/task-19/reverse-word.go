package main

import "fmt"

/**
 * Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
 * Символы могут быть unicode.
 */

func reverse(s string) string {
	runeStr := []rune(s)
	for i := 0; i < (len(runeStr)-1)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-i-1] = runeStr[len(runeStr)-i-1], runeStr[i]
	}
	return string(runeStr)
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
