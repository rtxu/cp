// Time: O(n^2)
// Space: O(1)

func finalPrices(prices []int) []int {
    n := len(prices)
    result := make([]int, n)
    for i := 0; i < n; i++ {
        var j int
        for j = i+1; j < n && prices[j] > prices[i]; j++ {}
        var discount int
        if j == n {
            discount = 0
        } else {
            discount = prices[j]
        }
        result[i] = prices[i] - discount
    }
    return result
}
