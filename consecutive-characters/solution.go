func maxPower(s string) int {
    n := len(s)
    if n == 1 {
        return 1
    }
    current, count := s[0], 1
    result := 1
    for i := 1; i < n; i++ {
        if s[i] == current {
            count++
            if count > result {
                result = count
            }
        } else {
            count = 1
            current = s[i]
        }
    }
    return result
}
