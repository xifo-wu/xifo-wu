package main

import "fmt"

func generate(numRows int) [][]int {
	row1 := []int{1}
	row2 := []int{1, 1}

	if numRows == 1 {
		return [][]int{row1}
	}

	res := [][]int{row1, row2}
	if numRows == 2 {
		return res
	}

	for i := 3; i <= numRows; i++ {
		t := []int{1}
		for j := 1; j < i-1; j++ {
			c := res[i-2][j-1] + res[i-2][j]
			t = append(t, c)
		}
		t = append(t, 1)
		res = append(res, t)
	}

	return res
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}
