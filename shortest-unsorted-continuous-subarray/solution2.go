// Time: O(n)
// Space: O(1)

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func findUnsortedSubarray(A []int) int {
    n := len(A)
    begin, end := -1, -2
    leftMax, rightMin := A[0], A[n-1]
    for i := 0; i < n; i++ {
        leftMax = max(leftMax, A[i])
        if A[i] < leftMax {
            end = i
        }
        rightMin = min(rightMin, A[n-1-i])
        if A[n-1-i] > rightMin {
            begin = n-1-i
        }
    }
    return end - begin + 1
}
