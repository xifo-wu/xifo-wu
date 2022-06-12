package main

import (
	"fmt"
)

func main() {
	len := 10000
	var k, n int
	fmt.Scanln(&k)

	nums := make([]int, k)
	arr := make([]int, len)

	for i := 0; i < k; i++ {
		fmt.Scan(&n)
		nums[i] = n
		arr[n] = 1
	}

	for _, v := range nums {
		if arr[v] == 1 {
			for v != 1 {
				if v%2 != 0 {
					v = (3*v + 1)
				}

				v /= 2

				arr[v] = 0
			}
		}
	}

	flag := false
	for i := len - 1; i >= 0; i-- {
		if arr[i] == 1 {
			if flag {
				fmt.Print(" ")
			}
			fmt.Print(i)
			flag = true
		}
	}

	return
}
