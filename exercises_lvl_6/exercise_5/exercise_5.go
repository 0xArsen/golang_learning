package main

import (
	"fmt"
)

type square struct {
	sideSize float64
}

func (s square) area() float64 {
	return s.sideSize * s.sideSize
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		sideSize: 3.0,
	}
	c := circle{
		radius: 3.0,
	}
	info(s)
	info(c)
}

func info(s shape) {
	fmt.Println(s.area())
}
