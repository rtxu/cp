package main

import "fmt"

var (
	dir = [4][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
)

func findPath(x, y int, path, matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])

	if path[x][y] != 0 {
		return
	} else {
		path[x][y] = 1
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if nx >= 0 && nx < row && ny >= 0 && ny < column && matrix[nx][ny] > matrix[x][y] {
				findPath(nx, ny, path, matrix)
				if path[nx][ny]+1 > path[x][y] {
					path[x][y] = path[nx][ny] + 1
				}
			}
		}
	}
}

func longestIncreasingPath(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	column := len(matrix[0])
	// path[i][j] is the longest increasing path begin at (i, j)
	path := make([][]int, row)
	for i := 0; i < row; i++ {
		path[i] = make([]int, column)
	}

	longestPath := 1
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			findPath(i, j, path, matrix)
			if path[i][j] > longestPath {
				longestPath = path[i][j]
			}
		}
	}
	return longestPath
}

func main() {
	var input [][]int
	var expected int
	input = [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	expected = 4
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, longestIncreasingPath(input))
	input = [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	expected = 4
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, longestIncreasingPath(input))
}
