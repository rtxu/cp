func jump(nums []int) int {
    n := len(nums)
    canReach := make([]int, n)
    for i := 1; i < n; i++ {
        canReach[i] = -1
    }
    
    maxReach := 0
    for i := 0; i < n; i++ {
        if canReach[i] >= 0 {
            var lastReach int
            for j := maxReach+1; j <= i+nums[i] && j < n; j++ {
                canReach[j] = canReach[i]+1
                lastReach = j
            }
            if lastReach > maxReach {
                maxReach = lastReach
            }
        }
    }
    
    return canReach[n-1]
}
