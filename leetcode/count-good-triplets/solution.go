func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
    n := len(arr)
    cnt := 0
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            for k := j+1; k < n; k++ {
                if abs(arr[i] - arr[j]) <= a && abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
                    cnt++
                }
            }
        }
    }
    return cnt
}
