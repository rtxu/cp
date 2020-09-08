func containsDuplicate(nums []int) bool {
    M := make(map[int]bool)
    n := len(nums)
    for i := 0; i < n; i++ {
        if M[nums[i]] {
            return true
        } else {
            M[nums[i]] = true
        }
    }
    return false
}
