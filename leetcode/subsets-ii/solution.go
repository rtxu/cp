func backtrack(start int, current, nums []int, result *[][]int) {
    currentCopy := make([]int, len(current))
    copy(currentCopy, current)
    *result = append(*result, currentCopy)
    for i := start; i < len(nums); i++ {
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        currentSnapshot := current
        current = append(current, nums[i])
        backtrack(i+1, current, nums, result)
        current = currentSnapshot
    }
}

func subsetsWithDup(nums []int) [][]int {
    var result [][]int
    sort.Ints(nums)
    backtrack(0, []int{}, nums, &result)
    return result
}
