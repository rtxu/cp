func countOdds(low int, high int) int {
    count := (high - low + 1) >> 1
    if low % 2 == 1 && high % 2 == 1 {
        count++
    }
    return count
}
