func combinationSum4(nums []int, target int) int {
    // A(i) 表示和为 i 的方案数, 0 <= i <= target
    // A(0) = 1
    // A(i) = Sum(A(i-k)), k in nums && i-k >= 0
    
    A := make([]int, target+1)
    A[0] = 1
    for i := 1; i <= target; i++ {
        var v int
        for k := 0; k < len(nums); k++ {
            if i - nums[k] >= 0 {
                v += A[i-nums[k]]
            }
        }
        A[i] = v
    }
    
    return A[target]
}
