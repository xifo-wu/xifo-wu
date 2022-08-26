package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return (nums[0] - 1) * (nums[1] - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}
