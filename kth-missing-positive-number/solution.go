func findKthPositive(arr []int, k int) int {
    n := len(arr)
    i, j := 1, 0
    for j < n && k > 0 {
        if arr[j] == i {
            j++
        } else {
            k--
        }
        i++
    }
    if j == n {
        return arr[j-1] + k
    } else {
        return i-1
    }
}
