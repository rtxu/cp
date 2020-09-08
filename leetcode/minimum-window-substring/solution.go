package main

import "fmt"

func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	count := make(map[byte]int)
	for i := 0; i < m; i++ {
		count[t[i]]++
	}
	nonZeroCount := len(count)

	var window string
	for i, j := 0, 0; j < n; j++ {
		if _, exist := count[s[j]]; exist {
			count[s[j]]--
			if count[s[j]] == 0 {
				nonZeroCount--
			}
			if nonZeroCount == 0 {
				for i < n {
					if _, exist := count[s[i]]; exist {
						if count[s[i]] == 0 {
							break
						} else {
							count[s[i]]++
							i++
						}
					} else {
						i++
					}
				}
				if len(window) == 0 || len(window) > j-i+1 {
					window = s[i : j+1]
				}
			}
		}
	}

	return window
}

func main() {
	var s, t, expected string
	s, t, expected = "awerwboiucbcwerabnc", "abc", "abnc"
	fmt.Printf("s: %s, t: %s, expected: %s, actual: %s\n", s, t, expected, minWindow(s, t))
	s, t, expected = "awerwboiucbcabnc", "abc", "bca"
	fmt.Printf("s: %s, t: %s, expected: %s, actual: %s\n", s, t, expected, minWindow(s, t))
	s, t, expected = "awerwboiubabn", "abc", ""
	fmt.Printf("s: %s, t: %s, expected: %s, actual: %s\n", s, t, expected, minWindow(s, t))
	s, t, expected = "awerwboiubabn", "", ""
	fmt.Printf("s: %s, t: %s, expected: %s, actual: %s\n", s, t, expected, minWindow(s, t))
	s, t, expected = "aa", "aa", "aa"
	fmt.Printf("s: %s, t: %s, expected: %s, actual: %s\n", s, t, expected, minWindow(s, t))
}
