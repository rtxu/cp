func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)
    if m * k > n {
        return -1
    }
    min, max := int(1e9+1), 0
    for i := 0; i < n; i++ {
        if bloomDay[i] > max {
            max = bloomDay[i]
        }
        if bloomDay[i] < min {
            min = bloomDay[i]
        }
    }
    
    possible := func (day int) bool {
        currentSum := 0
        flowerCnt := 0
        for i := 0; i < n; i++ {
            if day >= bloomDay[i] {
                currentSum++
                if currentSum == k {
                    flowerCnt++
                    currentSum = 0
                }
            } else {
                currentSum = 0
            }
        }
        return flowerCnt >= m
    }
    
    // left impossible, right possible
    left, right := min-1, max
    for left+1 != right {
        mid := (left+right)>>1
        if possible(mid) {
            right = mid
        } else {
            left = mid
        }
    }
    return right
}
