func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func longestPalindromeSubseq(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    var sb strings.Builder
    for i := 0; i < n; i++ {
        sb.WriteByte(s[n-1-i])
    }
    reverseS := sb.String()
    //fmt.Printf("origin = %s, reverse = %s\n", s, reverseS)
    
    // LCS(i, j) 表示 s 的 [0, i) 子串与 reverseS 的 [0, j) 子串的最长公共子串长度
    // LCS(0, j) = LCS(i, 0) = 0, 0 <= i, j <= n
    // LCS(n, n) 为最终结果
    LCS := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        LCS[i] = make([]int, n+1)
    }
    //fmt.Println(LCS[0])
    
    // 令字符串下标从 1 开始
    index := func(i int) int { return i-1 }
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            var v int
            if s[index(i)] == reverseS[index(j)] {
                v = max(v, LCS[i-1][j-1]+1)
            }
            v = max(v, LCS[i-1][j])
            v = max(v, LCS[i][j-1])
            LCS[i][j] = v
        }
        //fmt.Println(LCS[i])
    }
    return LCS[n][n]
}
