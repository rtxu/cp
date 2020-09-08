func max(a, b byte) byte {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b byte) byte {
    return a + b - max(a, b)
}

func makeGood(s string) string {
    n := len(s)
    var sb strings.Builder
    for i := 0; i < n-1; i++ {
        if max(s[i], s[i+1]) - min(s[i], s[i+1]) == 32 {
            return makeGood(s[:i] + s[i+2:])
        } else {
            sb.WriteByte(s[i])
        }
    }
    if n > 0 {
        sb.WriteByte(s[n-1])
    }
    return sb.String()
}
