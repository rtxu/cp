func index(i int) int {
    return i-1
}

func longestPrefix(s string) string {
    n := len(s)
    
    // 令 s 的小标从 1 开始，LPS[i] 表示 s[1...i] 已经匹配成功，i+1 位置发生不匹配，
    // 则已经匹配的 i 个字符中的最长后缀的长度，该后缀同时为 s 的前缀
    LPS := make([]int, n+1)
    LPS[1] = 0
    matched := 0
    for i := 2; i <= n; i++ {
        for matched > 0 && s[index(matched+1)] != s[index(i)] {
            matched = LPS[matched]
        }
        if s[index(matched+1)] == s[index(i)] {
            matched++
        }
        LPS[i] = matched
        fmt.Printf("LPS[%d] = %d\n", i, LPS[i])
    }
    fmt.Println("LPS: ", LPS[n])
    
    return s[:LPS[n]]
    
}
