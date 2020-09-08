func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func calculateMinimumHP(dungeon [][]int) int {
    // H(i, j) 为在 (i, j) 需要的最小健康值
    // H(i, j) = 
    // 1) dungeon[i][j] 为正，意味着该位置可以补充健康值，则 H(i, j) = min(H(i+1, j), H(i, j+1)) - dungeon[i][j]，
    //    但自身能量是不能出现 0 或负值的，故 H(i, j) = max(min(H(i+1, j), H(i, j+1)) - dungeon[i][j], 1)
    // 2) dungeon[i][j] 为负或0，意味着该位置需要消耗健康值，则 H(i, j) = min(H(i+1, j), H(i, j+1))-dungeon[i][j]
    // 特殊情况，H(row-1, col-1) 因为没有后继节点，其 minH = 1
    
    row := len(dungeon)
    if row == 0 {
        return -1
    }
    col := len(dungeon[0])
    
    H := make([][]int, row)
    for i := 0; i < row; i++ {
        H[i] = make([]int, col)
    }
    
    for i := row-1; i >= 0; i-- {
        for j := col-1; j >= 0; j-- {
            minH := -1
            if i+1 < row {
                minH = H[i+1][j]
            }
            if j+1 < col {
                if minH == -1 || H[i][j+1] < minH {
                    minH = H[i][j+1]
                }
            }
            if minH == -1 {
                minH = 1
            }
            if dungeon[i][j] > 0 {
                H[i][j] = max(minH-dungeon[i][j], 1)
            } else {
                H[i][j] = minH-dungeon[i][j]
            }
        }
    }
    
    return H[0][0]
}
