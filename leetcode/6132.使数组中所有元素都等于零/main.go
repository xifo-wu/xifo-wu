package main

import (
	"fmt"
	"sort"
)

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	r := 0
	d := 0
	for _, v := range nums {
		if v == 0 {
			continue
		}

		if v-d != 0 {
			r++
		}

		d = v

	}

	return r
}

func main() {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
	fmt.Println(minimumOperations([]int{0}))
}
