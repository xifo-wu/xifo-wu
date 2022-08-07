// https://leetcode.cn/contest/weekly-contest-305/problems/number-of-arithmetic-triplets/
package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
	length := len(nums)
	sum := 0
	for i := 0; i < length; i++ {
		curr := nums[i]
		currAddCount := 0
		for j := i + 1; j < length; j++ {
			if curr+diff == nums[j] {
				curr += diff
				currAddCount++
				if currAddCount == 2 {
					sum += 1
					break
				}
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
	fmt.Println(arithmeticTriplets([]int{0, 0}, 2))
}
