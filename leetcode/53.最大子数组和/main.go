package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums), 6)

	nums = []int{1}
	fmt.Println(maxSubArray(nums), 1)

	nums = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums), 23)
}
