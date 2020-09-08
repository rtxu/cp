func findDuplicate(nums []int) int {
    n := len(nums)
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    for i := 0; i < n; i++ {
        for nums[nums[i]-1] != nums[i] {
            swap(nums[i]-1, i)
        }
    }
    return nums[n-1]
}
