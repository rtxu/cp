var x = [4]int{0, 1, 0, -1}
var y = [4]int{1, 0, -1, 0}

func dfs(i, j int, grid [][]int, visited [][]bool) {
    n, m := len(grid), len(grid[0])
    if i < 0 || i >= n || j < 0 || j >= m {
        return
    }
    if grid[i][j] != 1 {
        return
    }
    if visited[i][j] {
        return
    }
    visited[i][j] = true
    for d := 0; d < 4; d++ {
        ni, nj := i + x[d], j + y[d]
        dfs(ni, nj, grid, visited)
    }
}

func count(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }
    cnt := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if !visited[i][j] && grid[i][j] == 1 {
                dfs(i, j, grid, visited)
                cnt++
            }
        }
    }
    // fmt.Println(cnt)
    return cnt
}

func minDays(grid [][]int) int {
    if count(grid) != 1 {
        return 0
    } else {
        n, m := len(grid), len(grid[0])
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                if grid[i][j] == 1 {
                    grid[i][j] = 0
                    if count(grid) != 1 {
                        return 1
                    }
                    grid[i][j] = 1
                }
            }
        }
    }
    return 2
}
