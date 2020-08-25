func minOperations(n int) int {
    cnt := 0
    median := n
    for i := 0; i < n / 2; i++ {
        cnt += median - (2 * i + 1)
    }
    return cnt
}
