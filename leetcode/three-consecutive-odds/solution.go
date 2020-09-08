func threeConsecutiveOdds(arr []int) bool {
    oddCnt := 0
    n := len(arr)
    for i := 0; i < n; i++ {
        if arr[i] % 2 == 1 {
            oddCnt++
            if oddCnt == 3 {
                return true
            }
        } else {
            oddCnt = 0
        }
    }
    return false
}

