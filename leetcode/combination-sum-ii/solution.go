func backtrack(i int, current []int, nums, counts []int, target int, result *[][]int) {
    if target == 0 {
        currentCopy := make([]int, len(current))
        copy(currentCopy, current)
        *result = append(*result, currentCopy)
        return
    }
    if i == len(nums) {
        return 
    }
    
    var j int
    for j = 0; j <= counts[i] && target - nums[i]*j >= 0; j++ {
        backtrack(i+1, current, nums, counts, target - nums[i]*j, result)
        current = append(current, nums[i])
    }
    current = current[0:len(current)-j]
}

func combinationSum2(candidates []int, target int) [][]int {
    M := make(map[int]int)
    for i := 0; i < len(candidates); i++ {
        M[candidates[i]]++
    }
    nums, counts := make([]int, len(M)), make([]int, len(M))
    var i int
    for num, count := range M {
        nums[i], counts[i] = num, count
        i++
    }
    var result [][]int
    backtrack(0, []int{}, nums, counts, target, &result)
    return result
}
