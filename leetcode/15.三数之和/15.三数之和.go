package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	r := [][]int{}

	for i := 0; i < l-1; i++ {
		left := i + 1
		right := l - 1

		//
		if i > 0 && nums[i] == nums[i-1] {
			fmt.Println(nums[i] == nums[i-1], i, nums[i], nums[i-1])
			continue
		}

		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				r = append(r, []int{nums[i], nums[left], nums[right]})
			}

			if sum > 0 {
				right--
			}

			if sum <= 0 {
				left++
			}
		}
	}

	return r
}

// @lc code=end

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0, 0}))
}
