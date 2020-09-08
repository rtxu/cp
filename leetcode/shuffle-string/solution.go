func restoreString(s string, indices []int) string {
    n := len(s)
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        result[indices[i]] = s[i]
    }
    return string(result)
}
