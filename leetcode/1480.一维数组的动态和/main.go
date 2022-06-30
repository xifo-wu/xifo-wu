package main

import "fmt"

func runningSum(nums []int) []int {
	sum := 0
	for k, v := range nums {
		sum += v
		nums[k] = sum
	}

	return nums
}

func main() {
	var s = []int{1, 2, 3, 4}
	fmt.Println(runningSum(s))
}
