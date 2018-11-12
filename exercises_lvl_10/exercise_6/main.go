package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	addToChan(c)
	pullFromChan(c)
}

func pullFromChan(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

func addToChan(c chan<- int) {
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
		}
		close(c)

	}()
}
