func longestSubarray(nums []int) int {
    n := len(nums)
    var zeroCnt, maxLen int
    for begin, end := 0, 0; end < n; end++ {
        if nums[end] == 0 {
            zeroCnt++
        }
        for zeroCnt > 1 {
            if nums[begin] == 0 {
                zeroCnt--
            }
            begin++
        }
        if end-begin > maxLen {
            maxLen = end-begin
        }
    }
    return maxLen
}
