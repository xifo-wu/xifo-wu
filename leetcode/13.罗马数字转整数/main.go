package main

import "fmt"

var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	r := len(s) - 1
	sum := 0
	p := 0
	for i := r; i >= 0; i-- {
		if m[s[i]] >= p {
			sum += m[s[i]]
		} else {
			sum -= m[s[i]]
		}

		p = m[s[i]]
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"), 3)
	fmt.Println(romanToInt("IV"), 4)
	fmt.Println(romanToInt("IX"), 9)
	fmt.Println(romanToInt("LVIII"), 58)
	fmt.Println(romanToInt("MCMXCIV"), 1994)
}
