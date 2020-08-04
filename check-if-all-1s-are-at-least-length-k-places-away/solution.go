func kLengthApart(nums []int, k int) bool {
    n := len(nums)
    last := -100000
    for i := 0; i < n; i++ {
        if nums[i] == 1 {
            distance := i - last
            if distance <= k {
                return false
            } 
            last = i
        }
    }
    return true
}
