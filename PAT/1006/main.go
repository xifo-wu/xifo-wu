package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	b := n / 100
	s := n / 10 % 10
	g := n % 10

	var result string
	for i := 0; i < b; i++ {
		result += "B"
	}

	for i := 0; i < s; i++ {
		result += "S"
	}

	for i := 0; i < g; i++ {
		result += strconv.Itoa(i + 1)
	}

	fmt.Println(result)
}
