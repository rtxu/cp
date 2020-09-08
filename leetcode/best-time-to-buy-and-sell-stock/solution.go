func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    
    maxProfit := 0
    minBuyPrice := prices[0]
    for i := 1; i < n; i++ {
        if prices[i] - minBuyPrice > maxProfit {
            maxProfit = prices[i] - minBuyPrice
        }
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        }
    }
    
    return maxProfit
}
