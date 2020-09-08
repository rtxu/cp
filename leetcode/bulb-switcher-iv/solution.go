func minFlips(target string) int {
    current := byte('0')
    n := len(target)
    flipCnt := 0
    for i := 0; i < n; i++ {
        if target[i] != current {
            flipCnt++
            current = target[i]
        }
    }
    return flipCnt
}
