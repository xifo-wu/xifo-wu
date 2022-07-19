package main

import "fmt"

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}

	return i
}

// @lc code=end

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement(a, 3), a)
	b := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(b, 2), b)
}
