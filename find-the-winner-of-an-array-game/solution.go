func getWinner(arr []int, k int) int {
    n := len(arr)
    winCnt := 0
    swap := func(i, j int) {
        arr[i], arr[j] = arr[j], arr[i]
    }
    for i := 1; i < n; i++ {
        if arr[i-1] > arr[i] {
            winCnt++
            swap(i-1, i)
        } else {
            winCnt = 1
        }
        arr = append(arr, arr[i-1])
        if winCnt == k {
            return arr[i]
        }
    }
    return arr[n-1]
}
