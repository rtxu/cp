// Space: O(1)

func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    isCol := false
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                if j == 0 {
                    isCol = true
                } else {
                    matrix[0][j] = 0
                }
            }
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if matrix[0][0] == 0 {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if isCol {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
