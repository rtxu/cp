func missingNumber(nums []int) int {
    n := len(nums)
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    for i := 0; i < n; i++ {
        for nums[i] < n && nums[nums[i]] != nums[i] {
            swap(nums[i], i)
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] == n {
            return i
        }
    }
    return n
}
