package main

import (
	"fmt"
	"strings"
)

func isTrue(s string) string {
	if len(s) < 3 {
		return "NO"
	}
	lp, pt, tr := 0, 0, 0
	pIndex, tIndex := 0, 0
	pCount, tCount := 0, 0
	for k, v := range s {
		tempS := string(v)
		// 出现别的字母就返回 NO
		if !(tempS == "P" || tempS == "A" || tempS == "T") {
			return "NO"
		}

		if tempS == string("P") {
			pCount += 1
			if pCount > 1 {
				return "NO"
			}

			pIndex = k
		}

		if tempS == "T" {
			tCount += 1
			if tCount > 1 {
				return "NO"
			}

			tIndex = k
		}
	}

	for k, v := range s {
		tempS := string(v)
		if k < pIndex && tempS == "A" {
			lp += 1
			continue
		}

		if k < tIndex && k > pIndex {
			pt += 1
			continue
		}

		if k > tIndex {
			tr += 1
			continue
		}
	}

	if pt == 0 {
		return "NO"
	}

	if lp*pt == tr {
		return "YES"
	}

	return "NO"
}

func main() {
	var n int
	fmt.Scanln(&n)
	i := 0
	a := make([]string, n)
	for i < n {
		var t string
		fmt.Scanln(&t)

		a[i] = isTrue(t)

		i += 1
	}

	fmt.Print(strings.Join(a, "\n"))
}
