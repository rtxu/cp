package main

import (
	"fmt"
	"strings"
)

var result []string
var words []string

func constructResult(n int, D []int, wordMap map[string]interface{}, s string) {
	// fmt.Printf("n = %d, D[n] = %d\n", n, D[n])
	if D[n] == 0 {
		return
	}
	if n == 0 {
		wordCount := len(words)
		for i := 0; i < wordCount/2; i++ {
			words[i], words[wordCount-1-i] = words[wordCount-1-i], words[i]
		}
		result = append(result, strings.Join(words, " "))
		for i := 0; i < wordCount/2; i++ {
			words[i], words[wordCount-1-i] = words[wordCount-1-i], words[i]
		}
		return
	}

	for word, _ := range wordMap {
		j := n - len(word)
		if j >= 0 && s[j:] == word {
			words = append(words, word)
			constructResult(j, D, wordMap, s[:j])
			words = words[:len(words)-1]
		}
	}
}

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]interface{})
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = nil
	}
	n := len(s)
	// D[i] means the number of valid sentences for s[:i]
	D := make([]int, n+1)
	D[0] = 1
	for i := 1; i <= n; i++ {
		for word, _ := range wordMap {
			j := i - len(word)
			if j >= 0 && s[j:i] == word {
				D[i] += D[j]
			}
		}
	}
	// fmt.Printf("the number of valid sentences: %d\n", D[n])

	result = make([]string, 0, D[n])
	words = make([]string, 0)
	constructResult(n, D, wordMap, s)
	return result
}

func main() {
	var s string
	var dict []string
	s, dict = "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Printf("s: %s, dict: %v, result: %v\n", s, dict, wordBreak(s, dict))

	s, dict = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("s: %s, dict: %v, result: %v\n", s, dict, wordBreak(s, dict))
	// too many answer
	/*
		s, dict = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
		fmt.Printf("s: %s, dict: %v, result: %v\n", s, dict, wordBreak(s, dict))
	*/
}
