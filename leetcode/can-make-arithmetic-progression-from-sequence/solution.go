func canMakeArithmeticProgression(arr []int) bool {
    n := len(arr)
    if n <= 1 {
        return true
    }
    sort.Ints(arr)
    diff := arr[0]-arr[1]
    for i := 2; i < n; i++ {
        if arr[i-1] - arr[i] != diff {
            return false
        }
    }
    return true
}
