func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func maxProfit(k int, prices []int) int {
    // A[i][j] 表示 0...i 完成 j 次交易的最优解
    // 0 <= i < n, 1 <= j <= k
    // 显然 A[i][1] 易计算
    // A[i][j] = max(A[i][j-1], A[i-1][j], t 的结果 + A[t-1][j-1]) 遍历最后一次计算发生的位置 t
    
    n := len(prices)
    if n < 2 {
        return 0
    }
    if k > n/2 {
        var result int
        for i := 1; i < n; i++ {
            if prices[i] > prices[i-1] {
                result += prices[i] - prices[i-1]
            }
        }
        return result
    }
    A := make([][]int, n)
    for i := 0; i < n; i++ {
        A[i] = make([]int, k+1)
    }
    
    for j := 1; j <= k; j++ {
        maxB := -prices[0]
        for i := 1; i < n; i++ {
            // 仅发生 j-1 次交易
            v := A[i][j-1]
            // 不在 i 位置发生交易
            if A[i-1][j] > v {
                v = A[i-1][j]
            }
            // 在 i 位置发生交易
            if prices[i] + maxB > v {
                v = prices[i] + maxB
            }
            A[i][j] = v
            maxB = max(maxB, -prices[i] + A[i-1][j-1])
        }
        if A[n-1][j] == A[n-1][j-1] {
            return A[n-1][j]
        }
        /*
        fmt.Println("j=", j)
        for i := 0; i < n; i++ {
            fmt.Printf("%d ", A[i][j])
        }
        fmt.Println()
        */
    }
    
    return A[n-1][k]
}
