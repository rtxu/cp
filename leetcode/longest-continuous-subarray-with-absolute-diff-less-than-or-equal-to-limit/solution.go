func longestSubarray(nums []int, limit int) int {
    n := len(nums)
    
    // minDeque 单调非递减，maxDeque 单调非递增
    minDeque, maxDeque := make([]int, 0), make([]int, 0)
    maxLen := 0
    for begin, end := 0, 0; end < n; end++ {
        current := nums[end]
        for len(minDeque) > 0 && minDeque[len(minDeque)-1] > current {
            minDeque = minDeque[:len(minDeque)-1]
        }
        minDeque = append(minDeque, current)
        for len(maxDeque) > 0 && maxDeque[len(maxDeque)-1] < current {
            maxDeque = maxDeque[:len(maxDeque)-1]
        }
        maxDeque = append(maxDeque, current)
        
        for begin <= end && maxDeque[0] - minDeque[0] > limit {
            popnum := nums[begin]
            if popnum == minDeque[0] {
                minDeque = minDeque[1:]
            }
            if popnum == maxDeque[0] {
                maxDeque = maxDeque[1:]
            }
            begin++
        }
        if end - begin + 1 > maxLen {
            maxLen = end - begin + 1
        }
        
    }
    return maxLen
}
