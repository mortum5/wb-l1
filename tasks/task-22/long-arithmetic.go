package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := value().SetString("190000050000000090000", 10)

	b, _ := value().SetString("155390000000111111000", 10)

	fmt.Println("Add: ", value().Add(a, b))
	fmt.Println("Sub: ", value().Sub(a, b))
	fmt.Println("Mul: ", value().Mul(a, b))
	fmt.Println("Div: ", value().Div(a, b))
}

func value() *big.Int {
	return new(big.Int)
}
