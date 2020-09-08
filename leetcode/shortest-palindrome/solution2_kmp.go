func index(i int) int { return i-1 }

func longestPrefixAsSuffix(s string) []int {
    n := len(s)
    // 设 s 的下标从 1 开始，LPS[i] 表示已经匹配了 [1...i] 个字符，第 i+1 个字符发生不匹配，此时与前缀相同的最长后缀的长度
    LPS := make([]int, n+1)
    LPS[1] = 0
    matched := 0
    for i := 2; i <= n; i++ {
        for matched > 0 && s[index(matched + 1)] != s[index(i)] {
            matched = LPS[matched]
        }
        if s[index(matched + 1)] == s[index(i)] {
            matched++
        }
        LPS[i] = matched
    }
    return LPS
}

func shortestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    
    reverseS := make([]byte, n)
    for i := 0; i < n; i++ {
        reverseS[i] = s[n-1-i]
    }
    temp := s + "$" + string(reverseS)
    LPS := longestPrefixAsSuffix(temp)
    fmt.Println(LPS[len(temp)])
    
    rightPadding := s[LPS[len(temp)]:]
    leftPadding := make([]byte, len(rightPadding))
    for i := 0; i < len(rightPadding); i++ {
        leftPadding[i] = rightPadding[len(rightPadding) - 1 -i]
    }
    return string(leftPadding) + s
}
