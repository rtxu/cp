func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func maxProfit(prices []int) int {
    // 某一段只能发生一次交易，所以如果能知道分段位置，答案立得
    // A[i] 表示截止到 i 位置获取到的最大利润
    // 则 A[0] = 0
    // A[1] = 是否能够和 0 构成交易？
    // 1) 如果能，则 A[1] = prices[1] - prices[0]
    // 2) 如果不能，0
    // A[i] = max(A[i-1], (prices[i] - prices[k]) + A[k-1]) 遍历之前可以构成交易的任意点 k
    
    // 令 B[i] = -prices[i] + A[i-1]
    n := len(prices)
    A := make([]int, n)
    A[0] = 0
    maxB := -prices[0]
    for i := 1; i < n; i++ {
        A[i] = max(A[i-1], prices[i] + maxB)
        maxB = max(maxB, -prices[i] + A[i-1])
    }
    
    return A[n-1]
}
