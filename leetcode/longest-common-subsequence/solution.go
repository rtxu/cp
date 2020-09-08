func longestCommonSubsequence(text1 string, text2 string) int {
    n1 := len(text1)
    n2 := len(text2)
    
    // LCS(i, j) 表示 s 的 [0, i) 子串与 reverseS 的 [0, j) 子串的最长公共子串长度
    // LCS(0, j) = LCS(i, 0) = 0, 0 <= i <= n1, 0 <= j <= n2
    // LCS(n1, n2) 为最终结果
    LCS := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        LCS[i] = make([]int, n2+1)
    }
    //fmt.Println(LCS[0])
    
    // 令字符串下标从 1 开始
    index := func(i int) int { return i-1 }
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            var v int
            if text1[index(i)] == text2[index(j)] {
                v = max(v, LCS[i-1][j-1]+1)
            }
            v = max(v, LCS[i-1][j])
            v = max(v, LCS[i][j-1])
            LCS[i][j] = v
        }
        //fmt.Println(LCS[i])
    }
    return LCS[n1][n2]
}

func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

