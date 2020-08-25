func maxDistance(position []int, m int) int {
    n := len(position)
    sort.Ints(position)
    left, right := 1, position[n-1] - position[0] + 1
    
    valid := func(dist int) bool {
        placedCnt := 0
        lastPos := -dist
        for i := 0; i < n; i++ {
            if position[i] - lastPos >= dist {
                placedCnt++
                lastPos = position[i]
            }
        }
        return placedCnt >= m
    }
    
    // left valid, right invalid
    for left + 1 < right {
        mid := (left + right) >> 1
        if valid(mid) {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}
