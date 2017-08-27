package main

import (
	"fmt"
)

func largestRectangleArea(h []int) int {
	n := len(h)
	if n == 0 {
		return 0
	}

	// lessFromLeft[i] means the first pos on the left side of i, meet h[lessFromLeft[i]] < h[i]
	lessFromLeft := make([]int, n)
	lessFromRight := make([]int, n)
	lessFromLeft[0] = -1
	lessFromRight[n-1] = n
	for i := 1; i < n; i++ {
		j := i - 1
		for j >= 0 && h[j] >= h[i] {
			j = lessFromLeft[j]
		}
		lessFromLeft[i] = j
	}
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && h[j] >= h[i] {
			j = lessFromRight[j]
		}
		lessFromRight[i] = j
	}
	var maxArea int
	for i := 0; i < n; i++ {
		v := (lessFromRight[i] - lessFromLeft[i] - 1) * h[i]
		if v > maxArea {
			maxArea = v
		}
	}
	return maxArea
}

func main() {
	var h []int
	var expected int
	h, expected = []int{2, 1, 5, 6, 2, 3}, 10
	fmt.Printf("h: %v, expected: %d, actual: %d\n", h, expected, largestRectangleArea(h))
	h, expected = []int{}, 0
	fmt.Printf("h: %v, expected: %d, actual: %d\n", h, expected, largestRectangleArea(h))
}
