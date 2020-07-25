
func minNumberOperations(target []int) int {
    n := len(target)
    steps := 0
    currentH := 0
    for i := 0; i < n; {
        maxH := -1
        for ; i < n && target[i] >= maxH; i++ { maxH = target[i] }
        // target[i] < maxH
        for ; i < n && target[i] <= target[i-1]; i++ {}
        // target[i] > target[i-1]
        steps += maxH - currentH
        currentH = target[i-1]
    }
    return steps
}

