func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func mergeStones(stones []int, K int) int {
    n := len(stones)
    if (n-1) % (K-1) > 0 {
        return -1
    }
    
    // intuition: 一个 sum 一定来自相邻的数
    // A(i, j, t) 表示 [i, j) 的数 merge 成 t 个数的最优值，由于 i, j 确定，则 t 确定
    // 故 A(i, j) 表示 [i, j) merge 后的最优解
    // 则 A(0, n) 为最终解
    // A(i, j) = min{
    //  A(i, t) + A(t, j) + sum(i, j) if pileSum = K
    //  A(i, t) + A(t, j) if pileSum < K
    // }, i < t < j
    // 令 pileCnt(i, j) 表示 [i, j) merge 后剩余的堆数
    // pileSum = pileCnt(i, t) + pileCnt(t, j)
    // 当 pileSum > K 时，这种情况不合法，考虑 K = 3, pileCnt(i, t) = pileCnt(t, j) = 2 的情况
    // 我们可以通过调整 t 获得 pileSum = 1 的合法情况
    // 对于所有的 j-i < K，merge cost = 0
    
    prefixSum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + stones[i]
    }
    
    pileCnt := func(n int) int {
        return (n-1) % (K-1) + 1
    }
    
    A := make([][]int, n)
    for i := 0; i < n; i++ {
        A[i] = make([]int, n+1)
    }
    for l := K; l <= n; l++ {
        for i := 0; i + l <= n; i++ {
            j := i+l
            A[i][j] = math.MaxInt32
            for t := i+1; t < j; t++ {
                pileSum := pileCnt(t-i) + pileCnt(j-t)
                if pileSum < K {
                    A[i][j] = min(A[i][j], A[i][t] + A[t][j])
                } else if pileSum == K {
                    A[i][j] = min(A[i][j], A[i][t] + A[t][j] + prefixSum[j] - prefixSum[i])
                }
            }
        }
    }
    
    return A[0][n]
}
