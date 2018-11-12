package main

import (
	"fmt"
)

func main() {
	fooResult := foo()
	barInt, barString := bar()

	fmt.Println(fooResult)
	fmt.Println(barInt, barString)

}

func foo() int {
	x := 42
	return x
}

func bar() (int, string) {
	x := 1
	y := "Hello"
	return x, y
}
