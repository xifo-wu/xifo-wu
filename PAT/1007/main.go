package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n > 2 && n%2 == 0 {
		return false
	}

	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}

	sum := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == 2 {
			sum += 1
		}
	}

	fmt.Println(sum)
}
