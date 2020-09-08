var x = [4]int{0, 1, 0, -1}
var y = [4]int{1, 0, -1, 0}

func findCycle(grid [][]byte, depth [][]int, parentDepth, i, j int) bool {
    m, n := len(grid), len(grid[0])
    if depth[i][j] > 0 {
        // visited
        return parentDepth - depth[i][j] + 1 >= 4
    } else {
        depth[i][j] = parentDepth + 1
        for d := 0; d < 4; d++ {
            ni, nj := i + x[d], j + y[d]
            if ni >= 0 && ni < m && nj >= 0 && nj < n && 
                grid[i][j] == grid[ni][nj] && 
                findCycle(grid, depth, depth[i][j], ni, nj) {
                return true
            }
        }
    }
    return false
}

func containsCycle(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    depth := make([][]int, m)
    for i := 0; i < m; i++ {
        depth[i] = make([]int, n)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if depth[i][j] != 0 {
                continue
            }
            if findCycle(grid, depth, 0, i, j) {
                return true
            }
        }
    }
    return false
}
