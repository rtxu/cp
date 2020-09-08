func kthFactor(n int, k int) int {
    j := 0
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            j++
        }
        if j == k {
            return i
        }
    }
    return -1
}
