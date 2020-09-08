func backtrack(nums []int, current []int, visited []bool, result *[][]int) {
    if len(current) == len(nums) {
        currentCopy := make([]int, len(current))
        copy(currentCopy, current)
        *result = append(*result, currentCopy)
        return
    }
    for i := 0; i < len(nums); i++ {
        if !visited[i] {
            visited[i] = true
            current = append(current, nums[i])
            backtrack(nums, current, visited, result)
            visited[i] = false
            current = current[0:len(current)-1]
        }
    }
}

func permute(nums []int) [][]int {
    visited := make([]bool,len(nums))
    var result [][]int
    backtrack(nums, []int{}, visited, &result)
    return result
}
