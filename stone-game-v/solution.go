func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    return a + b - min(a, b)
}

func umax(current *int, val int) {
    if val > *current {
        *current = val
    }
}

func stoneGameV(stoneValue []int) int {
    // T(i, j) 表示 [i, j) 的最优值，则 T(0, n) 为最终解
    // T(i, j) = max{min(Sum(i, k), Sum(k, j)) + T(min_pile)}, i < k < j
    // 如果 Sum(i, k) == Sum(k, j)
    n := len(stoneValue)
    prefixSum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + stoneValue[i]
    }
    // [i, j)
    rangeSum := func(i, j int) int {
        return prefixSum[j] - prefixSum[i]
    }
    T := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        T[i] = make([]int, n+1)
        for j := 0; j <= n; j++ {
            T[i][j] = -1
        }
    }
    for i := 0; i < n; i++ {
        T[i][i+1] = 0
    }
    for l := 2; l <= n; l++ {
        for i := 0; i+l <= n; i++ {
            j := i+l
            v := &T[i][j]
            for k := i; k < j; k++ {
                sum1, sum2 := rangeSum(i, k), rangeSum(k, j)
                current := min(sum1, sum2)
                if sum1 < sum2 {
                    current += T[i][k]
                } else if sum1 > sum2 {
                    current += T[k][j]
                } else {
                    current += max(T[i][k], T[k][j])
                }
                umax(v, current)
            }
        }
    }
    return T[0][n]
}
