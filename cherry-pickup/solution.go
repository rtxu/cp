func max(a []int) int {
    result := -1
    for i := 0; i < len(a); i++ {
        if a[i] > result {
            result = a[i]
        }
    }
    return result
}

// ref: https://leetcode.com/problems/cherry-pickup/discuss/165218/Java-O(N3)-DP-solution-w-specific-explanation
func cherryPickup(grid [][]int) int {
    // P(x1, y1, x2, y2) 表示 A 走到 (x1, y1) B 走到 (x2, y2) 时的最大值
    // 利用恒等式 x1+y1 = x2+y2，对状态进行压缩，P(x1, y1, x2)，隐含 y2 = x1+y1-x2
    // 为了方便边界条件的处理，这里假设 grid 范围从 1 到 n，0 行 0 列的 P 值均为 0
    // 对于任意状态 P(x1, y1, x2) = 
    // 1) 如果 grid(x1, y1) 或 grid(x2, y2) 为 -1，则当前状态为 -1
    // 2) 否则，P = 前一状态的最优解 + 当前状态可获得的额外 cherry 数
    // 前一状态的最优解 = 
    // 1) (A right, B right), P(x1, y1-1, x2)
    // 2) (A right, B down ), P(x1, y1-1, x2-1)
    // 3) (A down , B right), P(x1-1, y1, x2)
    // 4) (A down , B down ), P(x1-1, y1, x2-1)
    // 当前状态可获得的额外 cherry 数：如果 (x1, y1) = (x2, y2) 当前仅能获得 grid(x1, y1) 颗草莓，否则 grid(x1, y1) + grid(x2, y2)

    n := len(grid)
    P := make([][][]int, n+1)
    for i := 0; i <= n; i++ {
        P[i] = make([][]int, n+1)
        for j := 0; j <= n; j++ {
            P[i][j] = make([]int, n+1)
        }
    }
    
    for x1 := 0; x1 <= n; x1++ {
        for y1 := 0; y1 <= n; y1++ {
            for x2 := 0; x2 <= n; x2++ {
                P[x1][y1][x2] = -1
            }
        }
    }
    
    P[0][1][0] = 0
    P[0][1][1] = 0
    P[1][0][0] = 0
    P[1][0][1] = 0
    for x1 := 1; x1 <= n; x1++ {
        for y1 := 1; y1 <= n; y1++ {
            for x2 := 1; x2 <= n; x2++ {
                y2 := x1+y1-x2
                if y2 >= 1 && y2 <= n {
                    if grid[x1-1][y1-1] == -1 || grid[x2-1][y2-1] == -1 {
                        P[x1][y1][x2] = -1
                    } else {
                        lastMax := max([]int{P[x1][y1-1][x2], P[x1][y1-1][x2-1], P[x1-1][y1][x2], P[x1-1][y1][x2-1]})
                        if lastMax == -1 {
                            P[x1][y1][x2] = -1
                        } else {
                            extra := grid[x1-1][y1-1] + grid[x2-1][y2-1]
                            if x1 == x2 && y1 == y2 {
                                extra = grid[x1-1][y1-1]
                            }
                            P[x1][y1][x2] = lastMax + extra
                        }
                    }
                }
            }
        }
    }
    
    if P[n][n][n] == -1 {
        return 0
    } else {
        return P[n][n][n]
    }
}
