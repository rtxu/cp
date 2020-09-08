func rangeSum(nums []int, n int, left int, right int) int {
    allSum := make([]int, n*(n+1)/2)
    allSumTail := 0
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            allSum[allSumTail] = sum
            allSumTail++
        }
    }
    sort.Ints(allSum)
    var result int
    const MOD = 1e9+7
    for i := left; i <= right; i++ {
        result += allSum[i-1]
        result %= MOD
    }
    return result
}
