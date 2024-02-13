package main

import (
	"fmt"
	"strings"
)

/*
 * К каким негативным последствиям может привести данный фрагмент кода, и как
 * это исправить? Приведите корректный пример реализации.
 *
 * var justString string
 *
 * func someFunc() {
 * 	v := createHugeString(1 << 10)
 * 	justString = v[:100]
 * }
 *
 * func main() {
 * 	someFunc()
 * }
 *
 */

// ~~~~~~~~~~~~~~~~~~~~~~~~~~
//
// 1. Код может долго работать, если создание большой строки реализовано
//    не через stringBuilder
// 2. Из-за наличия глобальной переменной в которую сохраняется слайс с указателем
//    на большой кусок выделенной памяти будет утечка памяти, т.к. GC не сможет
//    очистить память из-за наличия ссылки
// 3. Не учитываются специфика unicode-символов, для чего надо конвертировать всё
//    через rune
//

const (
	N int = 1 << 10
)

var justString string

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func createHugeString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("ü")
	}
	return sb.String()
}

func someFunc() {
	v := createHugeString(N)

	n := min(N, 100)
	dst := make([]rune, n)

	copy(dst, ([]rune(v))[:n])

	justString = string(dst)
}

func main() {
	someFunc()

	fmt.Print(justString)
}
