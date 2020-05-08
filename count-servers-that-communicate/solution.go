func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    rowSum := make([]int, m)
    colSum := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            rowSum[i] += grid[i][j]
            colSum[j] += grid[i][j]
        }
    }
    
    nServers := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if rowSum[i] > 1 || colSum[j] > 1 {
                    nServers++
                }
            }
        }
    }
    return nServers
}
