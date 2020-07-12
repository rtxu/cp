const MOD = 1e9+7

func numSub(s string) int {
    n := len(s)
    oneCount := 0
    var result int
    for begin, end := 0, 0; end < n; end++ {
        if s[end] == '1' {
            oneCount++
        }
        for begin <= end && end-begin+1 != oneCount {
            if s[begin] == '1' {
                oneCount--
            }
            begin++
        }
        result += end-begin+1
        result %= MOD
    }
    return result
}
