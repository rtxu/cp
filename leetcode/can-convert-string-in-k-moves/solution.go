func canConvertString(s string, t string, k int) bool {
    ns, nt := len(s), len(t)
    if ns != nt {
        return false
    }
    n := ns
    var move [26]int
    for i := 0; i < n; i++ {
        if s[i] != t[i] {
            j := int((t[i] - s[i] + 26) % 26)
            if move[j] == 0 {
                move[j] = j
            } else {
                move[j] += 26
            }
            if move[j] > k {
                return false
            }
        }
    }
    return true
}
