func reformat(s string) string {
    n := len(s)
    d := make([]byte, 0)
    ch := make([]byte, 0)
    for i := 0; i < n; i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            ch = append(ch, s[i])
        } else {
            d = append(d, s[i])
        }
    }
    nd := len(d)
    nch := len(ch)
    if nd < nch {
        d, ch = ch, d
        nd, nch = nch, nd
    }
    if nd - nch <= 1 {
        var result strings.Builder
        i, j := 0, 0
        for ; i < nd && j < nch; i, j = i+1, j+1 {
            result.WriteByte(d[i])
            result.WriteByte(ch[j])
        }
        for ; i < nd; i++ {
            result.WriteByte(d[i])
        }
        return result.String()
    } else {
        return ""
    }
}
