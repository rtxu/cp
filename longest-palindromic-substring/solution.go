package main

import "fmt"

func longestPalindrome(s string) string {
	length, substring := 0, ""
	n := len(s)
	for i := 0; i < n-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if len(s[l+1:r]) > length {
			length = len(s[l+1 : r])
			substring = s[l+1 : r]
		}
	}
	for i := 0; i < n; i++ {
		l, r := i-1, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if len(s[l+1:r]) > length {
			length = len(s[l+1 : r])
			substring = s[l+1 : r]
		}
	}
	return substring
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("abcdef"))
	fmt.Println(longestPalindrome("bbbb"))
}
