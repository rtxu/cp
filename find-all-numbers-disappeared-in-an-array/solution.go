func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    swap := func (i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    for i := 0; i < n; i++ {
        for nums[nums[i]-1] != nums[i] {
            swap(nums[i]-1, i)
        }
    }
    var result []int
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            result = append(result, i+1)
        }
    }
    return result
}
