func ways(pizza []string, k int) int {
    // A(x1, x2, y1, y2, k) 考虑到 x2, y2 永远不变 => A(x, y, k)
    row, col := len(pizza), len(pizza[0])
    A := make([][][]int, row+1)
    for x := 0; x <= row; x++ {
        A[x] = make([][]int, col+1)
        for y := 0; y <= col; y++ {
            A[x][y] = make([]int, k+1)
            for i := 0; i <= k; i++ {
                A[x][y][i] = -1
            }
        }
    }
    prefixSum := make([][]int, row+1)
    for i := 0; i <= row; i++ {
        prefixSum[i] = make([]int, col+1)
    }
    pizzaCount := func(i, j int) int {
        if pizza[i][j] == 'A' {
            return 1
        } else {
            return 0
        }
    }
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            prefixSum[i+1][j+1] = prefixSum[i][j+1] + prefixSum[i+1][j] - prefixSum[i][j] + pizzaCount(i, j)
        }
    }
    hasApple := func(x1, x2, y1, y2 int) bool {
        count := prefixSum[x2][y2] - prefixSum[x1][y2] - prefixSum[x2][y1] + prefixSum[x1][y1]
        return count > 0
    }
    
    return solve(0, 0, row, col, k, A, hasApple)
}

const MOD = 1e9+7

func solve(x, y, row, col int, k int, A [][][]int, hasApple func(int, int, int, int) bool) int {
    count := A[x][y][k]
    if count >= 0 {
        return count
    }
    
    if k == 1 {
        if hasApple(x, row, y, col) {
            return 1
        } else {
            return 0
        }
    }
    
    count = 0
    for h := x + 1; h < row; h++ {
        if hasApple(x, h, y, col) {
            count = (count + solve(h, y, row, col, k-1, A, hasApple)) % MOD
        }
    }
    
    for v := y + 1; v < col; v++ {
        if hasApple(x, row, y, v) {
            count = (count + solve(x, v, row, col, k-1, A, hasApple)) % MOD
        }
    }
    
    A[x][y][k] = count
    return count
}
