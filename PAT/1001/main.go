package main

import "fmt"

func main() {
	var n, r int
	fmt.Scanln(&n)

	for n != 1 {
		r += 1
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3*n + 1) / 2
		}
		if n == 1 {
			break
		}
	}

	fmt.Println(r)
}
