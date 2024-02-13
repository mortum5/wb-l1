package main

import (
	"fmt"
	"strconv"
)

/**
 * Реализовать паттерн «адаптер» на любом примере.
 */

// Byte type alias for print override
type Byte byte

func (b Byte) String() string {
	return strconv.FormatInt(int64(b), 2)
}

// Binary And implementation
type And struct{}

func (And) Apply(a, b Byte) Byte {
	return a & b
}

// Unary Not implementation
type Not struct{}

func (Not) Apply(a Byte) Byte {
	return ^a
}

type UnaryOperation interface {
	Apply(A Byte) Byte
}

// Adapter that allows use all Unary operation as it was Binary operation
type UnaryAdapter struct {
	unaryOperation UnaryOperation
}

func (u UnaryAdapter) Apply(a, b Byte) Byte {
	return u.unaryOperation.Apply(a)
}

type BinaryOperation interface {
	Apply(A, B Byte) Byte
}

// Simple function that run any binary operation with arguments
func Calculate(operation BinaryOperation, a, b Byte) Byte {
	return operation.Apply(a, b)
}

func main() {
	a := Byte(131)
	b := Byte(31)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(Calculate(And{}, a, b))

	fmt.Println(Calculate(
		UnaryAdapter{Not{}},
		a,
		b,
	))
}
