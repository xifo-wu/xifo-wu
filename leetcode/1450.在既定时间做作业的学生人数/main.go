package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	sum := 0
	for k, v := range startTime {
		if queryTime >= v && queryTime <= endTime[k] {
			sum++
		}
	}

	return sum
}

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}
