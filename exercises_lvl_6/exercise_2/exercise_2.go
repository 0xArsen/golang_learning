package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fooResult := foo(nums...)
	barResult := bar(nums)

	fmt.Println(fooResult)
	fmt.Println(barResult)

}

func foo(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum

}

func bar(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
