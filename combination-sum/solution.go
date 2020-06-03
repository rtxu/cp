func constructResult(count, candidates []int) []int {
    n := len(candidates)
    var result []int
    for i := 0; i < n; i++ {
        for j := 0; j < count[i]; j++ {
            result = append(result, candidates[i])
        }
    }
    return result
}

func backtrack(i, sum int, count, candidates []int, target int, result *[][]int) {
    if sum > target {
        return
    }
    // fmt.Println(i, sum, count)
    
    if sum == target {
        // construct result
        (*result) = append(*result, constructResult(count, candidates))
        return
    }
    if i == len(candidates) {
        return
    }
    
    for cnt := 0; sum+candidates[i]*cnt <= target; cnt++ {
        count[i] = cnt
        backtrack(i+1, sum + candidates[i]*cnt, count, candidates, target, result)
    }
    count[i] = 0
}

func combinationSum(candidates []int, target int) [][]int {
    n := len(candidates)
    count := make([]int, n)
    var result [][]int
    backtrack(0, 0, count, candidates, target, &result)
    return result
}
