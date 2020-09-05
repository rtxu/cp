const MOD = 1e9+7

func mmul(a, b int) int {
    return int(int64(a) * int64(b) % MOD)
}

func numWays(s string) int {
    n := len(s)
    oneCnt := 0
    begin, end := make([]int, n+1), make([]int, n+1)
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            oneCnt++
            begin[oneCnt] = i
            end[oneCnt] = i
        } else {
            end[oneCnt] = i
        }
    }
    if oneCnt == 0 {
        return int((int64(n-1) * int64(n-2) / 2) % MOD)
    } else if oneCnt % 3 != 0 {
        return 0
    } else {
        oneThird, towThird := oneCnt / 3, 2 * oneCnt / 3
        return mmul((end[oneThird] - begin[oneThird] + 1), (end[towThird] - begin[towThird] + 1))
    }
}
