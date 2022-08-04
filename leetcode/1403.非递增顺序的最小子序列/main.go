package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1403 lang=golang
 *
 * [1403] 非递增顺序的最小子序列
 */

// @lc code=start
func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, v := range nums {
		total += v
	}

	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		if sum > total-sum {
			return nums[:i+1]
		}
	}
}

// @lc code=end

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))

}
