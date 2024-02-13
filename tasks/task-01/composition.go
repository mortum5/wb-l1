package main

import "fmt"

/**
 * Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
 * в структуре Action от родительской структуры Human (аналог наследования).
 */

type Head int
type Hand int

type Human struct {
	head      Head
	leftHand  Hand
	rightHand Hand
}

func (h Human) Drink() {
	fmt.Println("To Drink")
}

func (h Human) Eat() {
	fmt.Println("To Eat")
}

type Action struct {
	Human
}

func main() {
	human := Human{}
	human.Drink()
	human.Eat()

	action := Action{}
	action.Eat()
	action.Drink()
}
