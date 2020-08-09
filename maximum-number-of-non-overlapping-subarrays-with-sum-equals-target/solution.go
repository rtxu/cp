func maxNonOverlapping(nums []int, target int) int {
    n := len(nums)
    sum := 0
    prefixSum := map[int]bool{
        0: true,
    }
    var count int
    for i := 0; i < n; i++ {
        sum += nums[i]
        if prefixSum[sum - target] {
            count++
            sum = 0
            prefixSum = map[int]bool{
                0:true,
            }
        } else {
            prefixSum[sum] = true
        }
    }
    return count
}
