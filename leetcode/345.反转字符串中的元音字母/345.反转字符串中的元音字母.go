package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */
// @lc code=start
func reverseVowels(s string) string {
	y := "aeiouAEIOU"
	t := []byte(s)
	r := len(s) - 1
	l := 0
	for l < r {
		lok := strings.Contains(y, string(t[l]))
		rok := strings.Contains(y, string(t[r]))

		if lok && rok {
			t[l], t[r] = t[r], t[l]
			l++
			r--
		}

		if !lok {
			l++
		}

		if !rok {
			r--
		}
	}

	return string(t)
}

// @lc code=end

func main() {
	fmt.Println(reverseVowels("hello"), "holle")
	fmt.Println(reverseVowels("leetcode"), "leotcede")
}
