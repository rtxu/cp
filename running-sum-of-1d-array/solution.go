func runningSum(nums []int) []int {
    n := len(nums)
    sums := make([]int, n)
    sums[0] = nums[0]
    for i := 1; i < n; i++ {
        sums[i] = sums[i-1]+nums[i]
    }
    return sums
}
