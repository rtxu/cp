func umax(current *int, val int) {
    if val > *current {
        *current = val
    }
}

func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    sum := 0
    for i := 0; i < n; i++ {
        sum += cardPoints[i]
    }
    window := n - k
    windowSum := 0
    var i int
    for i = 0; i < window; i++ {
        windowSum += cardPoints[i]
    }
    var result int
    umax(&result, sum - windowSum)
    for ; i < n; i++ {
        windowSum = windowSum - cardPoints[i-window] + cardPoints[i]
        umax(&result, sum - windowSum)
    }
    return result
}
