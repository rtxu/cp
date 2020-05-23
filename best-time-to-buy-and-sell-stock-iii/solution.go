func maxProfit(prices []int) int {
    // 枚举两次交易的分割点
    n := len(prices)
    if n < 2 {
        return 0
    }
    maxProfit1 := make([]int, n)
    minBuyPrice := prices[0]
    for i := 1; i < n; i++ {
        maxProfit1[i] = maxProfit1[i-1]
        if prices[i] - minBuyPrice > maxProfit1[i] {
            maxProfit1[i] = prices[i] - minBuyPrice
        }
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        }
    }
    
    maxProfit2 := make([]int, n)
    maxSellPrice := prices[n-1]
    for i := n-2; i >= 0; i-- {
        maxProfit2[i] = maxProfit2[i+1]
        if maxSellPrice - prices[i] > maxProfit2[i] {
            maxProfit2[i] = maxSellPrice - prices[i]
        }
        if prices[i] > maxSellPrice {
            maxSellPrice = prices[i]
        }
    }
    
    var result int
    for i := 0; i < n; i++ {
        v := maxProfit1[i] 
        if i+1 < n {
            v += maxProfit2[i+1]
        }
        if v > result {
            result = v
        }
    }
    
    return result
}
