package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1374 lang=golang
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */

// @lc code=start
func generateTheString(n int) string {
	if n%2 == 0 {
		return strings.Repeat("x", n-1) + "y"
	}

	return strings.Repeat("x", n)
}

// @lc code=end

func main() {
	fmt.Println(generateTheString(4))
	fmt.Println(generateTheString(5))
}
