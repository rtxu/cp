func min(a ...int) int {
    v := a[0]
    for i := 1; i < len(a); i++ {
        if a[i] < v {
            v = a[i]
        }
    }
    return v
}

func minDifference(nums []int) int {
    n := len(nums)
    if n <= 4 {
        return 0
    }
    sort.Ints(nums)
    
    return min(nums[n-1] - nums[3], nums[n-2] - nums[2], nums[n-3] - nums[1], nums[n-4] - nums[0])
}
