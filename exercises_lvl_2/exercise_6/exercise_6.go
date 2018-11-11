package main

import (
	"fmt"
)

const (
	year        = 2018 + iota
	yearPlusOne = 2018 + iota
	yearPlusTwo = 2018 + iota
)

func main() {
	fmt.Println(year, yearPlusOne, yearPlusTwo)
}
