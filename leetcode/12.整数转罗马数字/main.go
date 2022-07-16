package main

import "fmt"

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, s := range valueSymbols {
		for num >= s.value {
			num -= s.value
			roman = append(roman, s.symbol...)
		}

		if num == 0 {
			break
		}
	}

	return string(roman)
}

func main() {
	fmt.Println(intToRoman(8))
}
