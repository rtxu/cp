// Time: O(n)
// Space: O(n)

func subarraySum(nums []int, k int) int {
    n := len(nums)
    preSum := map[int]int{0:1}
    
    var sum int
    var count int
    for i := 0; i < n; i++ {
        sum += nums[i]
        count += preSum[sum-k]
        preSum[sum]++
    }
    return count
}
