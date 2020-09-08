// Space: O(m+n)

func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    zeroRow := make([]bool, m)
    zeroCol := make([]bool, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                zeroRow[i] = true
                zeroCol[j] = true
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if zeroRow[i] || zeroCol[j] {
                matrix[i][j] = 0
            }
        }
    }
}
