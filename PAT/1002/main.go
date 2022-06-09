package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	fmt.Scanln(&n)
	enumMap := make(map[string]string)

	enumMap["1"] = "yi"
	enumMap["2"] = "er"
	enumMap["3"] = "san"
	enumMap["4"] = "si"
	enumMap["5"] = "wu"
	enumMap["6"] = "liu"
	enumMap["7"] = "qi"
	enumMap["8"] = "ba"
	enumMap["9"] = "jiu"
	enumMap["0"] = "ling"

	a := strings.Split(n, "")

	sum := 0
	for _, v := range a {
		t, _ := strconv.Atoi(v)

		sum += t
	}

	ra := strings.Split(strconv.Itoa(sum), "")

	rs := make([]string, len(ra))
	for i, v := range ra {
		rs[i] = enumMap[v]
	}

	fmt.Println(strings.Join(rs, " "))
}
