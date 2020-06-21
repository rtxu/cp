func xorOperation(n int, start int) int {
    var result int
    for i := 0; i < n; i++ {
        result ^= start + 2*i
    }
    return result
}
