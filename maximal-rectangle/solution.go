package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	var n, m int
	n = len(matrix)
	if n > 0 {
		m = len(matrix[0])
	}
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				sum[i][j] = 0
			} else {
				var delta int
				if matrix[i-1][j-1] == '0' {
					delta = 0
				} else {
					delta = 1
				}
				sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + delta
			}
		}
	}

	area := func(i1, j1, i2, j2 int) int {
		fullArea := (i2 - i1 + 1) * (j2 - j1 + 1)
		// fmt.Println("in area()", sum[i2+1][j2+1], sum[i2+1][j1], sum[i1][j2+1], sum[i1][j1])
		if sum[i2+1][j2+1]-sum[i2+1][j1]-sum[i1][j2+1]+sum[i1][j1] == fullArea {
			return fullArea
		} else {
			return -1
		}
	}
	// fmt.Println(area(1, 2, 2, 4))

	var maxArea int
	var max int
	for i1 := 0; i1 < n; i1++ {
		max = (n - i1) * m
		if max <= maxArea {
			break
		}
		for j1 := 0; j1 < m; j1++ {
			max = (n - i1) * (m - j1)
			if max <= maxArea {
				break
			}
			for i2 := n - 1; i2 >= i1; i2-- {
				max = (i2 - i1 + 1) * (m - j1)
				if max <= maxArea {
					break
				}
				for j2 := m - 1; j2 >= j1; j2-- {
					max = (i2 - i1 + 1) * (j2 - j1 + 1)
					if max <= maxArea {
						break
					}
					v := area(i1, j1, i2, j2)
					if v > maxArea {
						maxArea = v
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
