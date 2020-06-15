func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
    count := map[int]int{}
    for i := 0; i < len(arr); i++ {
        count[arr[i]]++
    }
    c := make([]int, len(count))
    i := 0
    for _, cnt := range count {
        c[i] = cnt
        i++
    }
    sort.Ints(c)
    removeCnt := 0
    for i := 0; i < len(c) && removeCnt < k; i++ {
        delta := min(k-removeCnt, c[i])
        c[i] -= delta
        removeCnt += delta
    }
    leftNumCount := 0
    for i := len(c)-1; i >= 0 && c[i] > 0; i-- {leftNumCount++}
    return leftNumCount
}
