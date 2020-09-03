func maxScore(s string) int {
    n := len(s)
    oneCnt := make([]int, n+1)
    for i := n-1; i >= 0; i-- {
        oneCnt[i] = oneCnt[i+1]
        if s[i] == '1' {
            oneCnt[i]++
        }
    }
    zeroCnt := 0
    var result int
    for i := 0; i < n-1; i++ {
        if s[i] == '0' {
            zeroCnt++
        }
        if zeroCnt + oneCnt[i+1] > result {
            result = zeroCnt + oneCnt[i+1]
        }
    }
    return result
}
