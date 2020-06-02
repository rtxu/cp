func canJump(nums []int) bool {
    n := len(nums)
    canReach := make([]bool, n)
    canReach[0] = true
    
    maxReach := 0
    for i := 0; i < n; i++ {
        if canReach[i] {
            var lastReach int
            for j := maxReach + 1; j <= i+nums[i] && j < n; j++ {
                canReach[j] = true
                lastReach = j
            }
            if lastReach > maxReach {
                maxReach = lastReach
            }
        }
    }
    
    return canReach[n-1]
}
