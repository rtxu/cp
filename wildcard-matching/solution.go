package main

import "fmt"

func isMatch(s string, p string) bool {
	n, m := len(p), len(s)

	/*
		D[i][j] = true, means p[:i] s[:j] is matched
		D[0][0] = true, initial state, D[n][m] is the result

		D[i][0] = true, if p[:i] = *... else false
		D[0][j] = false

		1 <= i <= n && 1 <= j <= m,
		switch p[i-1] {
		case letter:
			D[i][j] = D[i-1][j-1] && (p[i-1] == s[j-1])
		case '?':
			D[i][j] = D[i-1][j-1]
		case '*':
			D[i][j] = true if any D[i-1][k] = true, 0 <= k <= j
				      false oterwise
		}

	*/
	D := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]bool, m+1)
	}
	D[0][0] = true
	for i := 0; i < n && p[i] == '*'; i++ {
		D[i+1][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			switch p[i-1] {
			case '?':
				D[i][j] = D[i-1][j-1]
			case '*':
				for k := j; k >= 0; k-- {
					D[i][j] = (D[i][j] || D[i-1][k])
				}
			default:
				D[i][j] = (D[i-1][j-1] && p[i-1] == s[j-1])
			}
		}
	}

	return D[n][m]
}

func main() {
	var s, p string
	var expected bool

	s, p, expected = "aa", "a", false
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "aa", "aa", true
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "aaa", "aa", false
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "aa", "*", true
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "aa", "a*", true
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "ab", "?*", true
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
	s, p, expected = "aab", "c*a*b", false
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))

	s, p, expected = "bbbaaabaababbabbbaabababbbabababaabbaababbbabbbabb", "*b**b***baba***aaa*b***", false
	fmt.Printf("s: %s, p: %s, expected: %v, actual: %v\n", s, p, expected, isMatch(s, p))
}
