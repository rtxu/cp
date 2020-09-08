const MOD = 1e9+7

func numOfSubarrays(arr []int) int {
    currentOdd, currentEven := 0, 0
    totalOdd := 0
    n := len(arr)
    for i := 0; i < n; i++ {
        var nextOdd, nextEven int
        if arr[i] % 2 == 0 {
            nextOdd = currentOdd
            nextEven = currentEven + 1
        } else {
            nextOdd = currentEven + 1
            nextEven = currentOdd
        }
        currentOdd, currentEven = nextOdd, nextEven
        totalOdd += currentOdd
        totalOdd %= MOD
    }
    return totalOdd
}
