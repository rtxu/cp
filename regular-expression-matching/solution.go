package main

import "fmt"

func isMatch(s string, p string) bool {
	si, slen := 0, len(s)
	pi, plen := 0, len(p)

	for si < slen && pi < plen {
		switch p[pi] {
		case '.':
			if pi+1 < plen && p[pi+1] == '*' {
				pi++
			} else {
				si++
				pi++
			}
		case '*':
			ch := p[pi-1]
			if ch == '.' {
				for i := 0; i <= slen-si; i++ {
					if isMatch(s[si+i:], p[pi+1:]) {
						return true
					}
				}
			} else {
				if isMatch(s[si:], p[pi+1:]) {
					return true
				}
				for i := 1; i <= slen-si && s[si+i-1] == ch; i++ {
					if isMatch(s[si+i:], p[pi+1:]) {
						return true
					}
				}
			}
			return false
		default:
			if pi+1 < plen && p[pi+1] == '*' {
				pi++
			} else {
				if s[si] == p[pi] {
					si++
					pi++
				} else {
					return false
				}
			}
		}
	}

	if si == slen {
		for plen > pi && p[plen-1] == '*' {
			plen -= 2
		}
		return pi == plen
	} else {
		return false
	}
}

func main() {
	fmt.Println(isMatch("cd", "c*cd"), "true")
	fmt.Println(isMatch("aa", "a"), "false")
	fmt.Println(isMatch("aa", "aa"), "true")
	fmt.Println(isMatch("aaa", "aa"), "false")
	fmt.Println(isMatch("aa", "a*"), "true")
	fmt.Println(isMatch("aa", ".*"), "true")
	fmt.Println(isMatch("ab", ".*"), "true")
	fmt.Println(isMatch("aab", "c*a*b"), "true")
	fmt.Println(isMatch("abcd", "d*"), "false")
	fmt.Println(isMatch("a", "ab*"), "true")
	fmt.Println(isMatch("ab", ".*c"), "false")

}
