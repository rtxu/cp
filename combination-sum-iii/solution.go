func backtrack(num int, current []int, k, remain int, result *[][]int) {
    if len(current) == k {
        if remain == 0 {
            currentCopy := make([]int, len(current))
            copy(currentCopy, current)
            *result = append(*result, currentCopy)
        }
        return
    }
    
    for i := num; i < 10; i++ {
        current = append(current, i)
        backtrack(i+1, current, k, remain - i, result)
        current = current[0:len(current)-1]
    }
}

func combinationSum3(k int, n int) [][]int {
    var result [][]int
    backtrack(1, []int{}, k, n, &result)
    return result
}
