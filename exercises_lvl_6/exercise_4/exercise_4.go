package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and my age is %v\n", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Andrew",
		last:  "Kaufman",
		age:   34,
	}
	p.speak()

}
