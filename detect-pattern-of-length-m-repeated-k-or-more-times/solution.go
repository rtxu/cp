func count(pattern, arr []int) int {
    n := len(arr)
    cnt := 0
    i, j := 0, 0
    for i < n && arr[i] == pattern[j] {
        i++
        j++
        if j == len(pattern) {
            j = 0
            cnt++
        }
    }
    return cnt
}

func containsPattern(arr []int, m int, k int) bool {
    n := len(arr)
    for i := 0; i + m < n; i++ {
        pattern := arr[i:i+m]
        if count(pattern, arr[i+m:]) >= k-1 {
            return true
        }
    }
    return false
}
