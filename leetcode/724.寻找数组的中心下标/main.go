package main

import "fmt"

func pivotIndex(nums []int) int {
	total, sum := 0, 0
	for _, v := range nums {
		total += v
	}

	for k, v := range nums {
		if total-v == 2*sum {
			return k
		}
		sum += v
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums), "3")

	nums = []int{1, 2, 3}
	fmt.Println(pivotIndex(nums), "-1")

	nums = []int{2, 1, -1}
	fmt.Println(pivotIndex(nums), "0")
}
