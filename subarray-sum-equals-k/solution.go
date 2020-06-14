// Time: O(n^2)
// Space: O(n)

func subarraySum(nums []int, k int) int {
    n := len(nums)
    sums := make([]int, n+1)
    // sums[i] 表示 nums[0...i-1] 的和
    sums[0] = 0
    for i := 1; i <= n; i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }
    
    var count int
    for i := 0; i < n; i++ {
        for j := i+1; j <= n; j++ {
            if sums[j]-sums[i] == k {
                count++
            }
        }
    }
    return count
}
