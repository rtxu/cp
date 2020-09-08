func findLengthOfShortestSubarray(arr []int) int {
    n := len(arr)
    j := n-1
    for j > 0 && arr[j-1] <= arr[j] { j-- }
    if j == 0 {
        return 0
    }
    result := j
    for i := 0; i < n; i++ {
        if i == 0 || arr[i-1] <= arr[i] {
            for j < n && arr[j] < arr[i] { j++ }
            if j - i - 1 < result {
                result = j - i - 1
            }
        } else {
            break
        }
    }
    return result
}
