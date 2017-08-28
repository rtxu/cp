package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	var n, m int
	n = len(matrix)
	if n > 0 {
		m = len(matrix[0])
	}

	row := make([][]int, n)
	var maxArea int
	for i := 0; i < n; i++ {
		row[i] = make([]int, m)
		for j := 0; j < m; j++ {
			delta := int(matrix[i][j] - '0')
			if j == 0 {
				row[i][j] = delta
			} else {
				if delta == 1 {
					row[i][j] = row[i][j-1] + 1
				} else {
					row[i][j] = 0
				}
			}

			if row[i][j] > 0 {
				width, height := -1, -1
				for i2 := i; i2 >= 0; i2-- {
					if width == -1 || row[i2][j] < width {
						width = row[i2][j]
					}
					if width == 0 {
						break
					}
					height = i - i2 + 1
					if width*height > maxArea {
						maxArea = width * height
					}
				}
			}
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
}
