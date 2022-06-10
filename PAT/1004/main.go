package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var minName, minID string
	var maxName, maxID string
	minScore := 101
	maxScore := 0

	index := 0
	for index < n {
		var tempName, tempID string
		var tempScore int

		fmt.Scanf("%s %s %d", &tempName, &tempID, &tempScore)

		if tempScore >= maxScore {
			maxName = tempName
			maxID = tempID
			maxScore = tempScore
		}

		if tempScore < minScore {
			minName = tempName
			minID = tempID
			minScore = tempScore
		}

		index += 1
	}

	fmt.Printf("%s %s\n%s %s\n", maxName, maxID, minName, minID)
}
