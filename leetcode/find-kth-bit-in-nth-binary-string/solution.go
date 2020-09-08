func invert(s string) string {
    n := len(s)
    var sb strings.Builder
    for i := 0; i < n; i++ {
        if s[i] == '0' {
            sb.WriteByte('1')
        } else {
            sb.WriteByte('0')
        }
    }
    return sb.String()
}

func reverse(s string) string {
    n := len(s)
    var sb strings.Builder
    for i := n-1; i >= 0; i-- {
        sb.WriteByte(s[i])
    }
    return sb.String()
}

func build(n int) string {
    if n == 1 {
        return "0"
    } else {
        s := build(n-1)
        return s + "1" + reverse(invert(s))
    }
}

func findKthBit(n int, k int) byte {
    s := build(n)
    return s[k-1]
}
