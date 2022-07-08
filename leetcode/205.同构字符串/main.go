package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	sa := make([]int, len(s))
	for k, v := range s {
		_, ok := m[v]
		if !ok {
			m[v] = k
		}

		sa[k] = m[v]
	}

	m2 := make(map[rune]int)
	ta := make([]int, len(t))
	for k, v := range t {
		_, ok := m2[v]
		if !ok {
			m2[v] = k
		}

		ta[k] = m2[v]
	}

	for k, v := range sa {
		if v != ta[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
