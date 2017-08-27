package main

import (
	"fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	n := len(heights)
	// minHeight[i][j] means the minimum height between [i, j)
	// so area[i][j] = (j-i)*minHeight[i][j]
	minHeight := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		minHeight[i] = make([]int, n+1)
	}
	var maxArea int
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if j == i+1 {
				minHeight[i][j] = heights[i]
			} else {
				minHeight[i][j] = int(math.Min(float64(minHeight[i][j-1]), float64(heights[j-1])))
			}
			if (j-i)*minHeight[i][j] > maxArea {
				maxArea = (j - i) * minHeight[i][j]
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
