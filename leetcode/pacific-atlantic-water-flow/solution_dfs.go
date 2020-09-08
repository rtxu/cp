
func dfs(matrix, flowed [][]int, i, j int, preHeight int, tag int) {
    m, n := len(matrix), len(matrix[0])
    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }
    if flowed[i][j] & tag == tag {
        return
    }
    if matrix[i][j] < preHeight {
        return
    }
    flowed[i][j] |= tag
    dfs(matrix, flowed, i+1, j, matrix[i][j], tag)
    dfs(matrix, flowed, i-1, j, matrix[i][j], tag)
    dfs(matrix, flowed, i, j+1, matrix[i][j], tag)
    dfs(matrix, flowed, i, j-1, matrix[i][j], tag)
}

func pacificAtlantic(matrix [][]int) [][]int {
    // 1 to pacific, 2 to atlantic, 3 to both
    m := len(matrix)
    if m == 0 {
        return [][]int{}
    }
    n := len(matrix[0])
    
    flowed := make([][]int, m)
    for i := 0; i < m; i++ {
        flowed[i] = make([]int, n)
    }
    
    for i := 0; i < m; i++ {
        dfs(matrix, flowed, i, 0, math.MinInt32, 1)
        dfs(matrix, flowed, i, n-1, math.MinInt32, 2)
    }
    for j := 0; j < n; j++ {
        dfs(matrix, flowed, 0, j, math.MinInt32, 1)
        dfs(matrix, flowed, m-1, j, math.MinInt32, 2)
    }

    var result [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if flowed[i][j] == 3 {
                result = append(result, []int{i, j})
            }
        }
    }
    return result
}
