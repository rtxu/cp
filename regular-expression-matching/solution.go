package main

import "fmt"

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)

	//fmt.Println(s, p)
	if plen == 0 {
		return slen == 0
	}
	if slen == 0 {
		if plen > 1 && p[1] == '*' {
			return isMatch(s, p[2:])
		} else {
			return false
		}
	}

	if (plen > 1 && p[1] != '*') || plen <= 1 {
		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	} else {
		for i := 0; i < slen && (s[i] == p[0] || p[0] == '.'); i++ {
			if isMatch(s[i+1:], p[2:]) {
				return true
			}
		}
		return isMatch(s, p[2:])
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
