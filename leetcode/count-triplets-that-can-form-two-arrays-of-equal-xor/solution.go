func countTriplets(arr []int) int {
    count := 0
    n := len(arr)
    for i := 0; i < n; i++ {
        a := 0
        for j := i+1; j < n; j++ {
            a ^= arr[j-1]
            b := 0
            for k := j; k < n; k++ {
                b ^= arr[k]
                if a == b {
                    count++
                }
            }
        }
    }
    return count
}
