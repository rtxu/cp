package main

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	// d[i][j] means the minimum distance between word1[:i] and word2[:j], 0 <= i <= n && 0 <= j <= m
	d := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([]int, m+1)
	}

	for j := 0; j <= m; j++ {
		d[0][j] = j // insert j characters
	}
	for i := 0; i <= n; i++ {
		d[i][0] = i // delete i characters
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var delta int
			if word1[i-1] == word2[j-1] {
				delta = 0
			} else {
				delta = 1
			}
			d[i][j] = int(math.Min(float64(d[i-1][j]+1), // insert the last characters
				math.Min(float64(d[i][j-1]+1), // delete the last characters
					float64(d[i-1][j-1]+delta)))) // replace the last characters
		}
	}

	return d[n][m]
}

func main() {
	var word1, word2 string
	var expected int
	word1, word2, expected = "", "", 0
	fmt.Printf("word1: %s, word2: %s, expected: %d, actual: %d\n", word1, word2, expected, minDistance(word1, word2))
	word1, word2, expected = "abc", "", 3
	fmt.Printf("word1: %s, word2: %s, expected: %d, actual: %d\n", word1, word2, expected, minDistance(word1, word2))
	word1, word2, expected = "", "abc", 3
	fmt.Printf("word1: %s, word2: %s, expected: %d, actual: %d\n", word1, word2, expected, minDistance(word1, word2))
	word1, word2, expected = "abd", "abc", 1
	fmt.Printf("word1: %s, word2: %s, expected: %d, actual: %d\n", word1, word2, expected, minDistance(word1, word2))
}
