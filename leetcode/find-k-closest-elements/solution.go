func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func findClosestElements(arr []int, k int, x int) []int {
    minDiff := 0
    minDiffStart := 0
    for i := 0; i < k; i++ {
        minDiff += abs(arr[i] - x)
    }
    currentDiff := minDiff
    for i := k; i < len(arr); i++ {
        currentDiff = currentDiff - abs(arr[i-k]-x) + abs(arr[i]-x)
        // fmt.Println(currentDiff, minDiff)
        if currentDiff < minDiff {
            minDiff = currentDiff
            minDiffStart = i-k+1
        }
    }
    return arr[minDiffStart:minDiffStart+k]
}
