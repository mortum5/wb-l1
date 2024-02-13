package main

import "fmt"

/**
 * Разработать программу, которая в рантайме способна определить тип переменной:
 * int, string, bool, channel из переменной типа interface{}.
 */

func assert(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println(val, " is interger")
	case string:
		fmt.Println(val, " is string")
	case bool:
		fmt.Println(val, " is boolean")
	case chan int:
		fmt.Println(val, " is channel")
	}
}

func main() {
	var i int = 5
	var s string = "abc"
	var b bool = true
	var ch chan int = make(chan int)

	assert(i)
	assert(s)
	assert(b)
	assert(ch)
}
