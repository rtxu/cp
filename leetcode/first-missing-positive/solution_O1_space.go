// Time: O(n)
// Space: O(1)

func firstMissingPositive(nums []int) int {
    n := len(nums)
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    for i := 0; i < n; i++ {
        for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            swap(nums[i]-1, i)
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return n+1
}
