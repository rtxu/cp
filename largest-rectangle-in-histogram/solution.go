package main

import (
	"fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	n := len(heights)
	// minHeight[i][j] means the minimum height between [i, j)
	// so area[i][j] = (j-i)*minHeight[i][j]
	var maxArea int
	for i := 0; i < n; i++ {
		min := heights[i]
		for j := i + 1; j <= n; j++ {
			min = int(math.Min(float64(min), float64(heights[j-1])))
			if (j-i)*min > maxArea {
				maxArea = (j - i) * min
			}
		}
	}
	return maxArea
}

func main() {
	var h []int
	var expected int
	h, expected = []int{2, 1, 5, 6, 2, 3}, 10
	fmt.Printf("h: %v, expected: %d, actual: %d\n", h, expected, largestRectangleArea(h))
}
