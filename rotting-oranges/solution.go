type Node struct {
    i, j int
    rottenDay int
}

func orangesRotting(grid [][]int) int {
    n := len(grid)
    if n == 0 {
        return -1
    }
    m := len(grid[0])
    
    rotten := make([][]bool, n)
    for i := 0; i < n; i++ {
        rotten[i] = make([]bool, m)
    }
    
    Q := make([]Node, 0)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 2 {
                Q = append(Q, Node{i, j, 0})
                rotten[i][j] = true
            }
        }
    }
    
    x := [4]int{0, 1, 0, -1}
    y := [4]int{1, 0, -1, 0}
    var minRottenDay int
    for len(Q) > 0 {
        current := Q[0]
        if current.rottenDay > minRottenDay {
            minRottenDay = current.rottenDay
        }
        Q = Q[1:]
        for i := 0; i < 4; i++ {
            next_i := current.i + x[i]
            next_j := current.j + y[i]
            if next_i >= 0 && next_i < n && next_j >= 0 && next_j < m &&
                grid[next_i][next_j] == 1 && !rotten[next_i][next_j] {
                Q = append(Q, Node{next_i, next_j, current.rottenDay+1})
                rotten[next_i][next_j] = true
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 && !rotten[i][j] {
                return -1
            }
        }
    }
    return minRottenDay
}
