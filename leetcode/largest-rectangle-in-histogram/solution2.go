package main

import (
	"fmt"
)

func largestRectangleArea(h []int) int {
	h = append(h, -1)
	n := len(h)
	var maxArea int
	var stack []int
	for i := 0; i < n; {
		if len(stack) == 0 || h[stack[len(stack)-1]] <= h[i] {
			stack = append(stack, i)
			i++
		} else {
			height := h[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			var left int
			if len(stack) == 0 {
				left = -1
			} else {
				left = stack[len(stack)-1]
			}
			v := (i - left - 1) * height
			if v > maxArea {
				maxArea = v
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
	h, expected = []int{}, 0
	fmt.Printf("h: %v, expected: %d, actual: %d\n", h, expected, largestRectangleArea(h))
}
