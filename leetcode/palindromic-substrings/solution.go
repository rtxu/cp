
func countSubstrings(s string) int {
    n := len(s)
    count := n
    // P[i][j] = true 表示 s[i:j+1] 为回文串, i <= j
    P := make([][]bool, n)
    for i := 0; i < n; i++ {
        P[i] = make([]bool, n)
        P[i][i] = true
    }
    for length := 1; length <= n; length++ {
        for start := 0; start + length < n; start++ {
            end := start + length
            if s[start] == s[end] && (start+1 > end-1 || P[start+1][end-1]) {
                P[start][end] = true
                count++
            } else {
                P[start][end] = false
            }
        }
    }
    
    return count
}
