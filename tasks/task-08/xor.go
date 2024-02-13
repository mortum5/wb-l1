package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
 */

type Number int64

func (n Number) Show() string {
	return strconv.FormatInt(int64(n), 2)
}

func (n *Number) SetBit(pos int, v Number) {
	if v != 0 {
		v = 1
	}
	*n = (*n & ^(1 << pos)) | (v << pos)
}

func main() {
	var v Number = math.MaxInt64
	fmt.Println(v.Show())

	v.SetBit(0, 0)
	fmt.Println(v.Show())

}
