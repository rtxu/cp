package main

import "fmt"

func search(wordMap map[string]interface{}, s string) []string {
	n := len(s)
	result := make([]string, 0)
	for i := 1; i <= n; i++ {
		if _, exist := wordMap[s[:i]]; exist {
			if i == n {
				result = append(result, s)
			} else {
				if suffixes := search(wordMap, s[i:]); suffixes == nil {
				} else {
					for j := 0; j < len(suffixes); j++ {
						result = append(result, s[:i]+" "+suffixes[j])
					}
				}
			}
		}
	}
	return result
}

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]interface{})
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = nil
	}

	return search(wordMap, s)
}

func main() {
	var s string
	var dict []string
	s, dict = "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Printf("s: %s, dict: %v, result: %v\n", s, dict, wordBreak(s, dict))

	s, dict = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("s: %s, dict: %v, result: %v\n", s, dict, wordBreak(s, dict))
}
