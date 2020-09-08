func modifyString(s string) string {
    n := len(s)
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        if s[i] == '?' {
            for j := 0; j < 26; j++ {
                ch := byte('a' + j)
                if (i == 0 || result[i-1] != ch) && (i == n-1 || s[i+1] != ch) {
                    result[i] = ch
                    break
                }
            }
        } else {
            result[i] = s[i]
        }
    }
    return string(result)
}
