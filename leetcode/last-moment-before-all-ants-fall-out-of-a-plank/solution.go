func getLastMoment(n int, left []int, right []int) int {
    max := math.MinInt32
    for i := 0; i < len(left); i++ {
        left[i] = n-left[i]
        if n-left[i] > max {
            max = n-left[i]
        }
    }
    for i := 0; i < len(right); i++ {
        if n - right[i] > max {
            max = n - right[i]
        }
    }
    return max
}
