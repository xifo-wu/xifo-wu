package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	x := m % n
	arr = append(arr[(n-x):], arr[:(n-x)]...)

	fmt.Println(strings.Join(arr, " "))
}
