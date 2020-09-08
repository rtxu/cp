package main

import "fmt"

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

func maximalRectangle(matrix [][]byte) int {
	var n, m int
	n = len(matrix)
	if n > 0 {
		m = len(matrix[0])
	}

	h := make([]int, m)
	var maxArea int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			delta := int(matrix[i][j] - '0')
			if i == 0 {
				h[j] = delta
			} else {
				if delta == 1 {
					h[j] = h[j] + delta
				} else {
					h[j] = 0
				}
			}
		}
		v := largestRectangleArea(h)
		if v > maxArea {
			maxArea = v
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maximalRectangle([][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}))
	fmt.Println(maximalRectangle([][]byte{
		[]byte("10"),
		[]byte("10"),
	}))
}
