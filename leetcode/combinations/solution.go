func backtrack(start int, current []int, n, k int, result *[][]int) {
    if len(current) == k {
        currentCopy := make([]int, len(current))
        copy(currentCopy, current)
        *result = append(*result, currentCopy)
        return
    }
    if start > n {
        return
    }
    for i := start; i <= n; i++ {
        current = append(current, i) 
        backtrack(i+1, current, n, k, result)
        current = current[0:len(current)-1]
    }
}

func combine(n int, k int) [][]int {
    var result [][]int
    backtrack(1, []int{}, n, k, &result)
    return result
}
