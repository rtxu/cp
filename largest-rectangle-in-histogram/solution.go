package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	var maxArea int
	var left, right int
	left = -1
	for right = 1; right < n && heights[right] >= heights[0]; right++ {
	}
	maxArea = (right - left - 1) * heights[0]
	for i := 1; i < n; i++ {
		if heights[i] > heights[i-1] {
			left = i - 1
			for right = i + 1; right < n && heights[right] >= heights[i]; right++ {
			}
		} else if heights[i] < heights[i-1] {
			for ; left >= 0 && heights[left] >= heights[i]; left-- {
			}
			for ; right < n && heights[right] >= heights[i]; right++ {
			}
		} else {
			// do nothing
		}
		if (right-left-1)*heights[i] > maxArea {
			maxArea = (right - left - 1) * heights[i]
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
