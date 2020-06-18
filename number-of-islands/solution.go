
var x = [4]int{1, 0, -1, 0}
var y = [4]int{0, 1, 0, -1}

func dfs(i, j int, grid [][]byte, visited [][]bool) {
    visited[i][j] = true
    
    for d := 0; d < 4; d++ {
        next_i := i + x[d]
        next_j := j + y[d]
        if next_i >= 0 && next_i < len(grid) && next_j >= 0 && next_j < len(grid[0]) && 
            grid[next_i][next_j] == '1' && !visited[next_i][next_j] {
            dfs(next_i, next_j, grid, visited)
        }
    }
}

func numIslands(grid [][]byte) int {
    n := len(grid)
    if n == 0 {
        return 0
    }
    m := len(grid[0])
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }
    var count int
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                dfs(i, j, grid, visited)
                count++
            }
        }
    }
    return count
}
