func backtrack(total int, current string, leftCnt int, result *[]string) {
    if len(current) == total {
        *result = append(*result, current)
        return
    }
    
    if leftCnt < total/2 {
        backtrack(total, current+"(", leftCnt+1, result)
    }
    rightCnt := len(current) - leftCnt
    if leftCnt > rightCnt {
        backtrack(total, current+")", leftCnt, result)
    }
}

func generateParenthesis(n int) []string {
    var result []string
    backtrack(2*n, "", 0, &result)
    return result
}
