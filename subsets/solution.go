func backtrack(start int, current []int, nums []int, result *[][]int) {
    if start == len(nums) {
        currentCopy := make([]int, len(current))
        copy(currentCopy, current)
        *result = append(*result, currentCopy)
        return
    }
    backtrack(start+1, current, nums, result)
    current = append(current, nums[start])
    backtrack(start+1, current, nums, result)
    current = current[0:len(current)-1]
}

func subsets(nums []int) [][]int {
    var result [][]int
    backtrack(0, []int{}, nums, &result)
    return result
}
