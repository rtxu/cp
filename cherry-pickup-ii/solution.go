const INFINITY = 1e9

func cherryPickup(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    
    // A(c1, c2, r) 表示 robot1 在 (r, c1) robot2 在 (r, c2) 时的最优解
    // 0 <= c1, c2 < cols, 0 <= r < rows
    // 初始条件：A(0, cols-1, 0) = grid[0][0] + grid[0][cols-1]
    // 转移：A(c1, c2, r) = A(c1-1/c1/c1+1, c2-1/c2/c2+1, r-1) + gridCherryCount
    // gridCherryCount = grid[r][c1] + grid[r][c2] if c1 != c2 else grid[r][c1]
    // 1 <= r < rows
    
    A := make([][][]int, cols)
    for c1 := 0; c1 < cols; c1++ {
        A[c1] = make([][]int, cols)
        for c2 := 0; c2 < cols; c2++ {
            A[c1][c2] = make([]int, rows)
            for r := 0; r < rows; r++ {
                A[c1][c2][r] = -INFINITY
            }
        }
    }
    A[0][cols-1][0] = grid[0][0] + grid[0][cols-1]
    
    cherryCount := func(c1, c2, r int) int {
        if c1 == c2 {
            return grid[r][c1]
        } else {
            return grid[r][c1] + grid[r][c2]
        }
    }
    
    for r := 1; r < rows; r++ {
        for c1 := 0; c1 < cols; c1++ {
            for c2 := 0; c2 < cols; c2++ {
                lastMax := -int(INFINITY)
                for lastC1 := c1-1; lastC1 <= c1+1; lastC1++ {
                    if lastC1 >= 0 && lastC1 < cols {
                        for lastC2 := c2-1; lastC2 <= c2+1; lastC2++ {
                            if lastC2 >= 0 && lastC2 < cols {
                                if v := A[lastC1][lastC2][r-1]; v > lastMax {
                                    lastMax = v
                                }
                            }
                        }
                    }
                }
                if lastMax == -INFINITY {
                    A[c1][c2][r] = -int(INFINITY)
                } else {
                    A[c1][c2][r] = lastMax + cherryCount(c1, c2, r)
                }
            }
        }
    }
    
    var result int
    for c1 := 0; c1 < cols; c1++ {
        for c2 := 0; c2 < cols; c2++ {
            v := A[c1][c2][rows-1]
            if v > result {
                result = v
            }
        }
    }
    
    return result
}
