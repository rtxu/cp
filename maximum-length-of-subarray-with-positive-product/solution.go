func umax(current *int, val int) {
    if val > *current {
        *current = val
    }
}

func umin(current *int, val int) {
    if val < *current {
        *current = val
    }
}

func getMaxLen(nums []int) int {
    n := len(nums)
    // 0 pos, 1 neg
    last := [2]int{-1, n}
    prefix := 0
    
    maxLen := 0
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
        } else if nums[i] < 0 {
            prefix = 1 - prefix
        } else {
            last = [2]int{i, n}
            prefix = 0
        }
        umax(&maxLen, i - last[prefix])
        umin(&last[prefix], i)
    }
    return maxLen
}
